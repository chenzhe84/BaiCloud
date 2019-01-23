package main

import (
	"fmt"
	"github.com/chenzhe84/BaiCloud/metadata-service/app"
	pb "github.com/chenzhe84/BaiCloud/metadata-service/proto/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("metadata"),
	)

	service.Init()

	db, err := gorm.Open("mysql", "root:root@tcp(192.168.99.100:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&pb.App{})

	repo := &app.Repository{
		Db: db,
	}
	// err = repo.SaveApp(&pb.App{
	// 	Name: "test",
	// 	Text: "test2",
	// })
	app, err := repo.GetAppById("22d12587-1298-473f-b59e-5c8ce60a2051")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(app)
	}

	// handler := &app.Handler{
	// 	Repo: &app.Repository{
	// 		Db: db,
	// 	},
	// }

	// if err = pb.RegisterAppServiceHandler(service.Server(), handler); err != nil {
	// 	fmt.Println(err)
	// }

	// if err = service.Run(); err != nil {
	// 	fmt.Println(err)
	// }
}
