package init

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/common/gload"
)

func init() {
	viperConfig()
	InitMysql()
	InitRedis()
}

func viperConfig() {
	viper.SetConfigFile("../common/config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper读取文件错误.error:", err)
	}
	err = viper.Unmarshal(&gload.CONFIG)
	if err != nil {
		fmt.Printf("viper解码错误.error:", err)
	}
	fmt.Printf("viper文件内容:%v", gload.CONFIG)
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
		fmt.Printf("链接mysql出错.error:", err)
	}
}

func InitRedis() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Network: fmt.Sprintf("%s:%s", gload.CONFIG.Redis.Host, gload.CONFIG.Redis.Port),
		DB:      gload.CONFIG.Redis.DB,
	})
	err := rdb.FlushDB(ctx).Err()
	if err != nil {
		fmt.Printf("redis链接错误.error:", err)
	}

}
