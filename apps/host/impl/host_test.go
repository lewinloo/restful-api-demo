package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/lewinloo/restful-api-demo/apps/host"
	"github.com/lewinloo/restful-api-demo/apps/host/impl"
	"github.com/lewinloo/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
)

var (
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)

	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "cn-guangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {
	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
	if err != nil {
		panic(err)
	}
	// 需要初始化全局logger
	zap.DevelopmentSetup()

	service = impl.NewHostServiceImpl()
}
