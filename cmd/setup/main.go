package main

import (
	"github.com/thebluefowl/zion/model"
	"github.com/thebluefowl/zion/postgres"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Tenant{})
	db.AutoMigrate(&model.Subscriber{})
	db.AutoMigrate(&model.Notifier{})
}
