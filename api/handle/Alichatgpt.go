package handle

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"project/common/gload"
	"project/common/model"
	"strconv"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Input struct {
	Messages []Message `json:"messages"`
}

type Parameters struct {
	ResultFormat string `json:"result_format"`
}

type RequestBody struct {
	Model      string     `json:"model"`
	Input      Input      `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type QwenResponse struct {
	Output struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		TotalTokens         int `json:"total_tokens"`
		OutputTokens        int `json:"output_tokens"`
		InputTokens         int `json:"input_tokens"`
		PromptTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"prompt_tokens_details"`
	} `json:"output"`
	RequestID string `json:"request_id"`
}

func ApiGpt(task_no string, context string) {
	updates := model.Aliyungpt{Status: 2}
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构建请求体
	requestBody := RequestBody{
		// 此处以qwen-plus为例，可按需更换模型名称。模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
		Model: "qwen-plus",
		Input: Input{
			Messages: []Message{
				//{
				//	Role:    "system",
				//	Content: "You are a helpful assistant.",
				//},
				{
					Role:    "user",
					Content: context,
				},
			},
		},
		Parameters: Parameters{
			ResultFormat: "message",
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		updates.Status = 3
		updates.Errors = err.Error()
		log.Fatal(err)
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation", bytes.NewBuffer(jsonData))
	if err != nil {
		updates.Status = 3
		updates.Errors += err.Error()
		log.Fatal(err)
	}
	// 设置请求头
	// 若没有配置环境变量，请用百炼API Key将下行替换为：apiKey := "sk-xxx"
	req.Header.Set("Authorization", "Bearer "+gload.CONFIG.AliyunGpt.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		updates.Status = 3
		updates.Errors += err.Error()
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		updates.Status = 3
		updates.Errors += err.Error()
		log.Fatal(err)
	}

	// 定义变量接收解析后的 JSON
	var qwenResp QwenResponse

	// 解析 bodyText 到结构体
	if err := json.Unmarshal(bodyText, &qwenResp); err != nil {
		updates.Status = 3
		updates.Errors += "解析响应失败: " + err.Error()
		log.Println("解析响应失败:", err)
	} else {
		// 成功解析后可以获取到回答内容
		fmt.Println("模型回复内容：", qwenResp.Output.Choices[0].Message.Content)
		updates.Status = 2
		updates.Request = qwenResp.Output.Choices[0].Message.Content
	}
	gload.DB.Table("aliyunGpt").Where("task_no = ?", task_no).Updates(&updates)
}

type SendGptReq struct {
	Content string `json:"content" form:"content" binding:"required"`
}

func SendGpt(c *gin.Context) {
	var req SendGptReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
	}

	var task_no string
	task_no = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(rand.Intn(8999)+1000)
	fmt.Println("task_no:", task_no)

	data := model.Aliyungpt{
		TaskNo:    task_no,
		Content:   req.Content,
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	err = gload.DB.Table("aliyunGpt").Create(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	go ApiGpt(task_no, req.Content)
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"task_no": task_no,
		},
	})
	return
}

type GetGptDataReq struct {
	TaskNo string `json:"task_no" form:"task_no" binding:"required"`
}

func GetGptData(c *gin.Context) {
	var req GetGptDataReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
	}
	var data model.Aliyungpt
	err = gload.DB.Table("aliyunGpt").Where("task_no = ?", req.TaskNo).Limit(1).Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	if data.Id == 0 {
		c.JSON(200, gin.H{
			"code": 200,
			"data": gin.H{},
			"msg":  "内容为空",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data.Request,
		"msg":  "查询成功",
	})
	return
}

func FlowGpt(task_no string, question string) {
	ctx := context.Background()
	client := openai.NewClient(
		option.WithAPIKey(gload.CONFIG.AliyunGpt.ApiKey),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	filename := "./files/" + task_no + ".txt"
	file, err := os.Create(filename)
	defer file.Close()
	write := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
	}

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Seed:  openai.Int(0),
		Model: "qwen-plus",
	})

	// optionally, an accumulator helper can be used
	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if content, ok := acc.JustFinishedContent(); ok {
			println("Content stream finished:", content)
		}

		// if using tool calls
		if tool, ok := acc.JustFinishedToolCall(); ok {
			println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
		}

		if refusal, ok := acc.JustFinishedRefusal(); ok {
			println("Refusal stream finished:", refusal)
		}
		// it's best to use chunks after handling JustFinished events
		if len(chunk.Choices) > 0 {
			write.WriteString(chunk.Choices[0].Delta.Content)
		}
	}
	defer write.Flush()
	if stream.Err() != nil {
		panic(stream.Err())
	}
	write.WriteString("[end]")

	data, err := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(data)
	var dataStr string
	for scanner.Scan() {
		dataStr += scanner.Text()
	}

	updates := model.Aliyungpt{}
	updates.Status = 2
	updates.Request = dataStr
	gload.DB.Table("aliyunGpt").Where("task_no = ?", task_no).Updates(&updates)
}

func SendFlowGpt(c *gin.Context) {
	var req SendGptReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
	}

	var task_no string
	task_no = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(rand.Intn(8999)+1000)
	fmt.Println("task_no:", task_no)

	data := model.Aliyungpt{
		TaskNo:    task_no,
		Content:   req.Content,
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	err = gload.DB.Table("aliyunGpt").Create(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	go FlowGpt(task_no, req.Content)
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"task_no": task_no,
		},
	})
	return
}

func GetFlowGpt(c *gin.Context) {
	var req GetGptDataReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
	}
	var data model.Aliyungpt
	err = gload.DB.Table("aliyunGpt").Where("task_no = ?", req.TaskNo).Limit(1).Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	if data.Id > 0 {
		if data.Status == 2 {
			c.JSON(200, gin.H{
				"code": 200,
				"data": data.Request,
				"msg":  "查询成功",
			})
		}
	}

	filename := "./files/" + req.TaskNo + ".txt"
	fileData, err := os.Open(filename)
	defer fileData.Close()
	scanner := bufio.NewScanner(fileData)
	var dataStr string
	for scanner.Scan() {
		dataStr += scanner.Text()
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": dataStr,
		"msg":  "查询成功",
	})
	return
}
