package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"github.com/lewinloo/restful-api-demo/apps/host"
)

// 创建 host
func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(&ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	// 进行接口调用
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, ins)
}

// 查询 host
func (h *Handler) queryHost(c *gin.Context) {
	// 从 context 中获取参数
	req := host.NewQueryHostFromContext(c)

	fmt.Printf("page_number: %d, page_size: %d, keywords: %s\n", req.PageNumber, req.PageSize, req.Keywords)

	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

// 查询 host 详情
func (h *Handler) describeHost(c *gin.Context) {
	// 从 context 中获取参数
	req := host.NewIdRequestFromContext(c)

	fmt.Printf("id: %s\n", req.Id)

	ins, err := h.svc.DescribeHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, ins)
}

// 删除 host
func (h *Handler) deleteHost(c *gin.Context) {
	// 从 context 中获取参数
	req := host.NewIdRequestFromContext(c)

	fmt.Printf("id: %s\n", req.Id)

	ins, err := h.svc.DeleteHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, ins)
}

// 更新 put host
func (h *Handler) putHost(c *gin.Context) {
	// 从 context 中获取参数
	req := host.NewPutUpdateHostRequest(c.Param("id"))
	if err := c.ShouldBindJSON(&req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")

	ins, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, ins)
}

// 更新 patch host
func (h *Handler) patchHost(c *gin.Context) {
	// 从 context 中获取参数
	req := host.NewPatchUpdateHostRequest(c.Param("id"))
	if err := c.ShouldBindJSON(&req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")

	ins, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, ins)
}
