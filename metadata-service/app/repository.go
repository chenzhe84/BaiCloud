package app

import (
	pb "github.com/chenzhe84/BaiCloud/metadata-service/proto/app"
	"github.com/jinzhu/gorm"
)

type IRepository interface {
	SaveApp(*pb.App) error
	GetAppById(string) (*pb.App, error)
	SaveModule(*pb.Module) error
	GetModuleById(string) (*pb.Module, error)
	SaveForm(*pb.Form) error
	GetFormById(string) (*pb.Form, error)
}

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) SaveApp(app *pb.App) error {
	return r.Db.Save(app).Error
}

func (r *Repository) GetAppById(id string) (*pb.App, error) {
	var app pb.App
	if err := r.Db.First(&app, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *Repository) SaveModule(m *pb.Module) error {
	return r.Db.Save(m).Error
}

func (r *Repository) GetModuleById(id string) (*pb.Module, error) {
	var module pb.Module
	if err := r.Db.First(&module, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *Repository) SaveForm(f *pb.Form) error {
	return r.Db.Save(f).Error
}

func (r *Repository) GetFormById(id string) (*pb.Form, error) {
	var form pb.Form
	if err := r.Db.First(&form, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &form, nil
}
