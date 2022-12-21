package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lewinloo/restful-api-demo/apps"
	"github.com/lewinloo/restful-api-demo/apps/host"
	"github.com/lewinloo/restful-api-demo/conf"
)

// 接口实现的静态检查
var impl = &HostServiceImpl{}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子logger
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

func (h *HostServiceImpl) Config() {
	h.l = zap.L().Named("Host")
	h.db = conf.C().MySQL.GetDB()
}

func (h *HostServiceImpl) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryImpl(impl)
}
