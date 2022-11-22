package main

import (
	"github.com/thebluefowl/zion/internal/pg"
	"github.com/thebluefowl/zion/model"
)

func main() {
	db, err := pg.Connect()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Tenant{})
	db.AutoMigrate(&model.Subscriber{})
	db.AutoMigrate(&model.Notifier{})
}
