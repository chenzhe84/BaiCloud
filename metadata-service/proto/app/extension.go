package service_metadata_app

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *App) BeforeCreate(scope *gorm.Scope) error {
	uuid, _ := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
