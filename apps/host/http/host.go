package http

import (
	"fmt"
	"strconv"

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
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("page_number", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	kw := c.Query("keywords")
	req := host.NewQueryHostRequest()
	req.PageNumber = pageNumber
	req.PageSize = pageSize
	req.Keywords = kw

	fmt.Printf("page_number: %d, page_size: %d, keywords: %s\n", req.PageNumber, req.PageSize, req.Keywords)

	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}
