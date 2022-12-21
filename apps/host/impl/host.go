package impl

import (
	"context"

	"github.com/lewinloo/restful-api-demo/apps/host"
)

func (h *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	h.l.Debugf("create host %s", ins.Name)
	// h.l.With(logger.NewAny("request-id", "bfefbwiwdw")).Debug("create host...")
	// 校验数据的合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 注入默认值填充
	ins.InjectDefault()

	if err := h.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (h *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (h *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}

func (h *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (h *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
