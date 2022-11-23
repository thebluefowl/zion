package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/model"
	"github.com/thebluefowl/zion/service"
)

type TenantHandler struct {
	e             *echo.Echo
	tenantService *service.TenantService
}

func NewTenantHandler(e *echo.Echo, tenantService *service.TenantService) *TenantHandler {
	return &TenantHandler{e: e, tenantService: tenantService}
}

type CreateTenantRequest struct {
	Name string `json:"name"`
}

func (h *TenantHandler) Create(c echo.Context) error {
	request := &CreateTenantRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(400, err)
	}
	tenant, err := h.tenantService.Create(&model.Tenant{
		ID:   ksuid.New().String(),
		Name: request.Name,
	})
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, tenant)
}

func (h *TenantHandler) Get(c echo.Context) error {
	id := c.Param("id")
	tenant, err := h.tenantService.Get(id)
	if err != nil {
		return c.JSON(500, err)
	}
	if tenant == nil {
		return c.JSON(404, nil)
	}
	return c.JSON(200, tenant)
}

func (h *TenantHandler) AddRoutes() {
	h.e.POST("/tenants", h.Create)
	h.e.GET("/tenants/:id", h.Get)
}
