package handler

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/thebluefowl/zion/model"
	"github.com/thebluefowl/zion/service"
)

type NotifierHandler struct {
	e               *echo.Echo
	notifierService *service.NotifierService
}

func NewNotifierHandler(e *echo.Echo, notifierService *service.NotifierService) *NotifierHandler {
	return &NotifierHandler{e: e, notifierService: notifierService}
}

// Notifiers
type CreateNotifier struct {
	SubscriberID string          `param:"subscriber_id"`
	TenantID     string          `param:"tenant_id"`
	NotifierType string          `json:"notifier_type"`
	Config       json.RawMessage `json:"config"`
}

func (h *NotifierHandler) Create(c echo.Context) error {
	request := &CreateNotifier{}
	if err := c.Bind(request); err != nil {
		return c.JSON(400, err)
	}
	if err := h.notifierService.Create(&service.AddNotifierRequest{
		SubscriberID: request.SubscriberID,
		TenantID:     request.TenantID,
		NotifierType: model.NotifierType(request.NotifierType),
		Config:       request.Config,
	}); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "OK")
}

type FilterNotifierRequest struct {
	SubscriberID string `param:"subscriber_id"`
	TenantID     string `param:"tenant_id"`
	NotifierType string `query:"notifier_type"`
	IsActive     bool   `query:"is_active"`
}

func (h *NotifierHandler) Filter(c echo.Context) error {
	request := &FilterNotifierRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(400, err)
	}
	notifiers, err := h.notifierService.Filter(request.SubscriberID, request.TenantID, model.NotifierType(request.NotifierType), request.IsActive)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, notifiers)

}

func (h *NotifierHandler) AddRoutes() {
	h.e.POST("/tenants/:tenant_id/subscribers/:subscriber_id/notifiers", h.Create)
	h.e.GET("/tenants/:tenant_id/subscribers/:subscriber_id/notifiers", h.Filter)
}
