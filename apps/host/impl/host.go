package impl

import (
	"context"

	"github.com/infraboard/mcube/sqlbuilder"
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
	sqb := sqlbuilder.NewBuilder(QueryHostSQL)
	if req.Keywords != "" {
		sqb.Where("r.`name` LIKE ? OR r.description LIKE ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	sqb.Limit(req.Offset(), req.GetPageSize())
	querySql, args := sqb.Build()
	h.l.Debugf("query sql: %s, args: %v", querySql, args)

	// query stmt, 构造一个Prepare语句
	stmt, err := h.db.PrepareContext(ctx, querySql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := host.NewHostSet()
	// 遍历查询数据库的数据
	for rows.Next() {
		ins := host.NewHost()
		// 扫描字段到对象里
		if err := rows.Scan(&ins.Id, &ins.Vendor, &ins.Region,
			&ins.CreateAt, &ins.ExpireAt, &ins.Type, &ins.Name,
			&ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
			&ins.Account, &ins.PublicIP, &ins.PrivateIP,
			&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec,
			&ins.OSType, &ins.OSName, &ins.SerialNumber); err != nil {
			return nil, err
		}
		set.Add(ins)
	}

	// total 统计
	countSQL, args := sqb.BuildCount()
	h.l.Debugf("count sql: %s, args: %v", countSQL, args)
	cstmt, err := h.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer cstmt.Close()
	if err := cstmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}

	return set, nil
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
