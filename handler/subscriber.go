package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/thebluefowl/zion/service"
)

type Handler struct {
	subscriberService *service.SubscriberService
	e                 *echo.Echo
}

func NewHandler(e *echo.Echo, subscriberService *service.SubscriberService) *Handler {
	return &Handler{e: e, subscriberService: subscriberService}
}
