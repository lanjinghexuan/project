package gload

import (
	"gorm.io/gorm"
	"project/common/config"
)

var (
	DB     *gorm.DB
	CONFIG config.Config
)
