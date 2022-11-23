package handler

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/thebluefowl/zion/model"
	"github.com/thebluefowl/zion/service"
)

type SubscriberHandler struct {
	subscriberService *service.SubscriberService
	e                 *echo.Echo
}

func NewSubscriberHandler(
	e *echo.Echo,
	subscriberService *service.SubscriberService,
) *SubscriberHandler {
	return &SubscriberHandler{e: e, subscriberService: subscriberService}
}

type CreateSubscriberRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	TenantID string `param:"tenant_id"`
}

func (h *SubscriberHandler) Create(c echo.Context) error {
	request := &CreateSubscriberRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(400, err)
	}
	subscriber, err := h.subscriberService.Create(&service.CreateRequest{
		Name:     request.Name,
		Email:    request.Email,
		TenantID: request.TenantID,
	})
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, subscriber)
}

func (h *SubscriberHandler) Get(c echo.Context) error {
	id := c.Param("id")
	tenantID := c.Param("tenant_id")

	subscriber, err := h.subscriberService.GetSubscriber(tenantID, id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, subscriber)
}

// Notifiers
type CreateNotifier struct {
	SubscriberID string          `param:"subscriber_id"`
	TenantID     string          `param:"tenant_id"`
	NotifierType string          `json:"notifier_type"`
	Config       json.RawMessage `json:"config"`
}

func (h *SubscriberHandler) AddNotifier(c echo.Context) error {
	request := &CreateNotifier{}
	if err := c.Bind(request); err != nil {
		return c.JSON(400, err)
	}
	if err := h.subscriberService.AddNotifier(&service.AddNotifierRequest{
		SubscriberID: request.SubscriberID,
		TenantID:     request.TenantID,
		NotifierType: model.NotifierType(request.NotifierType),
		Config:       request.Config,
	}); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "OK")
}

func (h *SubscriberHandler) AddRoutes() {
	h.e.POST("/tenants/:tenant_id/subscribers", h.Create)
	h.e.GET("/tenants/:tenant_id/subscribers/:id", h.Get)
	h.e.POST("/tenants/:tenant_id/subscribers/:subscriber_id/notifiers", h.AddNotifier)
}
