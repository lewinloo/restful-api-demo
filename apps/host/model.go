package host

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type HostSet struct {
	Total int     `json:"total"`
	Items []*Host `json:"items"`
}

func (h *HostSet) Add(item *Host) {
	h.Items = append(h.Items, item)
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

// Host 模型定义
type Host struct {
	// 资源公共属性部分
	*Resource
	// 资源独有属性
	*Describe
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

type Vendor int

const (
	// 私有云
	PRIVATE_IDC Vendor = iota
	// 阿里云
	ALIYUN
	// 腾讯云
	TENCENTYUN
)

type Resource struct {
	Id          string            `json:"id" validate:"required"`     // 全局唯一ID
	Vendor      Vendor            `json:"vendor"`                     // 厂商
	Region      string            `json:"region" validate:"required"` // 地区
	CreateAt    int64             `json:"create_at"`                  // 创建时间
	ExpireAt    int64             `json:"expire_at"`                  // 过期时间
	Type        string            `json:"type" validate:"required"`   // 规格
	Name        string            `json:"name" validate:"required"`   // 名称
	Description string            `json:"description"`                // 描述
	Status      string            `json:"status"`                     // 状态
	Tags        map[string]string `json:"tags"`                       // 标签
	UpdateAt    int64             `json:"update_at"`                  // 更新时间
	SyncAt      int64             `json:"sync_at"`                    // 同步时间
	Account     string            `json:"account"`                    // 资源所属账号
	PublicIP    string            `json:"public_ip"`                  // 公网IP
	PrivateIP   string            `json:"private_ip"`                 // 内网IP
	// PayType     string            `json:"pay_type"`                   // 实例付费方式
}

type Describe struct {
	ResourceId   string `json:"resource_id"`                // 关联 Resource
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为windows、linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
	//ImageID                 string `json:"image_id"`                    // 镜像ID
	//InternetMaxBandWidthOut int    `json:"internet_max_band_width_out"` // 公网出带宽最大值，单位为Mbps
	//InternetMaxBandWidthIn  int    `json:"internet_max_band_width_in"`  // 公网入带宽最大值，单位为Mbps
	//KeyPairName             string `json:"key_pair_name"`               // 私钥对名称
	//SecurityGroups          string `json:"security_groups"`             // 安全组，采用逗号分割
}

func NewQueryHostFromContext(c *gin.Context) *QueryHostRequest {
	pageNumber := c.Query("page_number")
	pageSize := c.Query("page_size")
	kw := c.Query("keywords")
	req := NewQueryHostRequest()
	if pageNumber != "" {
		req.PageNumber, _ = strconv.Atoi(pageNumber)
	}
	if pageSize != "" {
		req.PageSize, _ = strconv.Atoi(pageSize)
	}
	req.Keywords = kw
	return req
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

type QueryHostRequest struct {
	PageSize   int    `json:"page_size" form:"page_size"`
	PageNumber int    `json:"page_number" form:"page_number"`
	Keywords   string `json:"keywords" form:"keywords"`
}

func (req *QueryHostRequest) Offset() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

func (req *QueryHostRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func NewIdRequestWithId(id string) *IdRequest {
	return &IdRequest{
		Id: id,
	}
}

func NewIdRequestFromContext(c *gin.Context) *IdRequest {
	id := c.Param("id")
	req := NewIdRequestWithId(id)
	return req
}

type IdRequest struct {
	Id string `json:"id"`
}

type UpdateHostRequest struct {
	*Describe
}
