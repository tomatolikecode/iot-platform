package svc

import (
	"github.com/iot-platform/admin/internal/config"
	"github.com/iot-platform/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
