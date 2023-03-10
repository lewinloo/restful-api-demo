package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lewinloo/restful-api-demo/apps"
	"github.com/lewinloo/restful-api-demo/apps/host"
)

var (
	handler = &Handler{}
)

type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	// 从 IOC 里面获取 HostService 对象
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}

func (h *Handler) Registry(r gin.IRouter) {
	host := r.Group("/hosts")
	{
		host.POST("", h.createHost)
		host.GET("", h.queryHost)
		host.GET("/:id", h.describeHost)
		host.DELETE("/:id", h.deleteHost)
		host.PUT("/:id", h.putHost)
		host.PATCH("/:id", h.patchHost)
	}
}

func (h *Handler) Name() string {
	return host.AppName
}

// 完成模块注册 http handler
func init() {
	apps.RegistryGin(handler)
}
