package pkr

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"time"
)

type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

func GetConfig() MinIOConfig {
	return MinIOConfig{
		Endpoint:  "127.0.0.1:9000",
		AccessKey: "ak",
		SecretKey: "sk",
		UseSSL:    false,
	}
}

func UploadFileToMinIO(bucketName, objectName string, file *multipart.FileHeader) (string, error) {
	//方法为实现参数自行编写
	//示例  devGetConfig
	config := GetConfig()
	addr := config.Endpoint
	// 初始化 MinIO 客户端
	minioClient, err := minio.New(addr, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return "", fmt.Errorf("failed to initialize MinIO client: %v", err)
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// 获取文件信息
	fileStat, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to get file stats: %v", err)
	}
	defer fileStat.Close()
	objectName = fmt.Sprintf("%s/%s", time.Now().Format("2006-01-02"), objectName)
	// 使用 PutObject 上传文件
	_, err = minioClient.PutObject(
		context.Background(),
		bucketName,
		objectName,
		src,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to MinIO: %v", err)
	}

	fmt.Printf("Successfully uploaded file %s to bucket %s as object %s\n",
		file.Filename, bucketName, objectName)
	return addr + "/" + bucketName + "/" + objectName, nil
}

/*

// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Failed to get uploaded file: " + err.Error(),
		})
		return
	}
	if file.Size > 1024*1024*50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件内容过大",
		})
	}

	// 验证文件扩展名
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".mp4":  true,
	}

	ext := strings.ToLower(path.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件类型不支持，仅支持 png, jpg, jpeg, mp4 格式",
		})
		return
	}

	// 直接上传到 MinIO
	fileUrl, err := middlewear.UploadFileToMinIO("lanjing", file.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to upload file to MinIO: " + err.Error(),
		})
		return
	}
*/
/*

// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Failed to get uploaded file: " + err.Error(),
		})
		return
	}
	if file.Size > 1024*1024*50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件内容过大",
		})
	}

	// 验证文件扩展名
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".mp4":  true,
	}

	ext := strings.ToLower(path.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件类型不支持，仅支持 png, jpg, jpeg, mp4 格式",
		})
		return
	}

	// 直接上传到 MinIO
	fileUrl, err := middlewear.UploadFileToMinIO("lanjing", file.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to upload file to MinIO: " + err.Error(),
		})
		return
	}
*/
