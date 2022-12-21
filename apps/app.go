package apps

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lewinloo/restful-api-demo/apps/host"
)

var (
	HostService host.Service

	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	implApps[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

func RegistryGin(svc GinService) {
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}
	ginApps[svc.Name()] = svc
}

// 初始化 impl 层
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}
	return names
}

// 初始化gin路由
func InitGin(r gin.IRouter) {
	for _, v := range ginApps {
		v.Config()
	}

	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

type GinService interface {
	Registry(r gin.IRouter)
	Name() string
	Config()
}
