package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thebluefowl/zion/handler"
	"github.com/thebluefowl/zion/internal/pg"
	"github.com/thebluefowl/zion/repository"
	"github.com/thebluefowl/zion/service"
)

func main() {
	db, err := pg.Connect()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	tenantRepository := repository.NewTenantRepository(db)
	tenantService := service.NewTenantService(tenantRepository)
	tenantHandler := handler.NewTenantHandler(e, tenantService)
	tenantHandler.AddRoutes()

	subscriberRepository := repository.NewSubscriberRepository(db)
	subscriberService := service.NewSubscriberService(subscriberRepository, tenantService)
	subscriberHandler := handler.NewSubscriberHandler(e, subscriberService)
	subscriberHandler.AddRoutes()

	e.Logger.Fatal(e.Start(":8080"))
}
