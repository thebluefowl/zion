package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thebluefowl/zion/postgres"
	"github.com/thebluefowl/zion/subscriber"
	"github.com/thebluefowl/zion/tenant"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	tenantRepository := tenant.NewTenantPGRepository(db)
	tenantService := tenant.NewService(tenantRepository)
	tenantHandler := tenant.NewHandler(e, tenantService)
	tenantHandler.AddRoutes()

	subscriberRepository := subscriber.NewSubscriberPGRepository(db)
	subscriberService := subscriber.NewService(subscriberRepository)
	subscriberHandler := subscriber.NewHandler(e, subscriberService, tenantService)
	subscriberHandler.AddRoutes()

	e.Logger.Fatal(e.Start(":8080"))

}
