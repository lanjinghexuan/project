package gload

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/lanjinghexuan/project/common/config"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Config
	ES     *elasticsearch.Client
	Ctx    = context.Background()
	MONGO  *mongo.Client
	REDIS  *redis.Client
)
