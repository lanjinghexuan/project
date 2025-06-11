package gload

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"project/common/config"
)

var (
	DB     *gorm.DB
	CONFIG config.Config
	ES     *elasticsearch.Client
	Ctx    = context.Background()
	MONGO  *mongo.Client
)
