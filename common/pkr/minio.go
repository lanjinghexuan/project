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
