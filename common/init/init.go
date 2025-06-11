package init

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"net/http"
	"project/common/gload"
	"time"
)

func init() {
	viperConfig()
	InitMysql()
	InitRedis()
	InitElasticsearch()
	InitMongo()
}

func viperConfig() {
	viper.SetConfigFile("../common/config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper读取文件错误.error:", err)
	}
	err = viper.Unmarshal(&gload.CONFIG)
	if err != nil {
		fmt.Println("viper解码错误.error:", err)
	}
	fmt.Println("viper文件内容:%v", gload.CONFIG)
}

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		gload.CONFIG.Mysql.UserName,
		gload.CONFIG.Mysql.Password,
		gload.CONFIG.Mysql.Host,
		gload.CONFIG.Mysql.Port,
		gload.CONFIG.Mysql.Database,
	)
	var err error
	gload.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("链接mysql出错.error:", err)
	}
	fmt.Println("mysql连接成功:%v", gload.DB)
}

func InitRedis() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Network: fmt.Sprintf("%s:%s", gload.CONFIG.Redis.Host, gload.CONFIG.Redis.Port),
		DB:      gload.CONFIG.Redis.DB,
	})
	err := rdb.FlushDB(ctx).Err()
	if err != nil {
		//fmt.Println("redis链接错误.error:", err)
	}

}

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%d", gload.CONFIG.Elastic.Host, gload.CONFIG.Elastic.Port),
		},
		Username: gload.CONFIG.Elastic.User,
		Password: gload.CONFIG.Elastic.Pass,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	var err error
	gload.ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("<UNK>elasticsearch<UNK>.error:", err)
	}
}

func InitMongo() {
	client := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d",
		gload.CONFIG.Mongo.User,
		gload.CONFIG.Mongo.Pass,
		gload.CONFIG.Mongo.Host,
		gload.CONFIG.Mongo.Port,
	))

	c, err := mongo.Connect(gload.Ctx, client)
	if err != nil {
		fmt.Printf("error connecting to database: %v\n", err)
		return
	}

	err = c.Ping(gload.Ctx, nil)
	if err != nil {
		fmt.Printf("error pinging database: %v\n", err)
		return
	}
	gload.MONGO = c
	fmt.Println(gload.MONGO)
}
