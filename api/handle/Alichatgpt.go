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
	fmt.Println("创建协程：" + task_no)
	ctx := context.Background()
	client := openai.NewClient(
		option.WithAPIKey(gload.CONFIG.AliyunGpt.ApiKey),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	filename := "./files/" + task_no + ".txt"

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close() // 确保文件在函数结束时关闭

	write := bufio.NewWriter(file)
	defer write.Flush() // 确保缓冲区内容被刷新到文件

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Seed:  openai.Int(0),
		Model: "qwen-plus",
	})

	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if content, ok := acc.JustFinishedContent(); ok {
			fmt.Println("Content stream finished:", content)
		}

		if tool, ok := acc.JustFinishedToolCall(); ok {
			fmt.Println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
		}

		if refusal, ok := acc.JustFinishedRefusal(); ok {
			fmt.Println("Refusal stream finished:", refusal)
		}

		if len(chunk.Choices) > 0 {
			if _, err := write.WriteString(chunk.Choices[0].Delta.Content); err != nil {
				fmt.Println("写入文件失败:", err)
				return
			}
			write.Flush()
			//fmt.Println("<UNK>", chunk.Choices[0].Delta.Content)
		}
	}

	if stream.Err() != nil {
		fmt.Println("流处理错误:", stream.Err())
		return
	}

	// 读取文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	dataStr := string(data)
	fmt.Println("文件内容:", dataStr)

	updates := model.Aliyungpt{}
	updates.Status = 2
	updates.Request = dataStr
	if err := gload.DB.Table("aliyunGpt").Where("task_no = ?", task_no).Updates(&updates).Error; err != nil {
		fmt.Println("更新数据库失败:", err)
		return
	}
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
	//通过task_no 查询
	err = gload.DB.Table("aliyunGpt").Where("task_no = ?", req.TaskNo).Limit(1).Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	if data.Id > 0 {
		if data.Status == 2 {
			//如果数据存在并且状态 == 2 返回数据库内容
			fmt.Println("查询到数据并且状态为2")
			c.JSON(200, gin.H{
				"code": 200,
				"data": data.Request,
				"msg":  "查询成功",
			})
			return
		}
	}
	fmt.Println("数据不存在或者状态不为2")
	filename := "./files/" + req.TaskNo + ".txt"
	fileData, err := os.Open(filename)
	if err != nil {
		fmt.Println("open文件错误:", err)
	}
	defer fileData.Close()
	scanner := bufio.NewScanner(fileData)

	var dataStr string
	for scanner.Scan() {
		fmt.Println("打开文件", scanner.Text())
		dataStr += scanner.Text()
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": dataStr,
		"msg":  "查询成功",
	})
	return
}
