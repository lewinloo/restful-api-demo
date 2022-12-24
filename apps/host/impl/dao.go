package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lewinloo/restful-api-demo/apps/host"
)

func (h *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	// 处理事务的提交方式
	// 无错误，则commit事务
	// 有错误，则rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				h.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				h.l.Error("commit error, %s", err)
			}
		}
	}()

	// 插入 resource 数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.ExecContext(ctx, ins.Id, ins.Vendor, ins.Region, ins.CreateAt,
		ins.ExpireAt, ins.Type, ins.Name, ins.Description, ins.Status,
		ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP, ins.PrivateIP)
	if err != nil {
		return err
	}

	// 插入 describe 数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.ExecContext(ctx, ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec, ins.OSType, ins.OSName, ins.SerialNumber)
	if err != nil {
		return err
	}

	return nil
}

func (h *HostServiceImpl) update(ctx context.Context, ins *host.Host) error {
	var (
		err      error
		resStmt  *sql.Stmt
		descStmt *sql.Stmt
	)

	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				h.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				h.l.Error("commit error, %s", err)
			}
		}
	}()

	// 更新 resource 表
	resStmt, err = tx.PrepareContext(ctx, UpdateResourceSQL)
	if err != nil {
		return err
	}
	defer resStmt.Close()
	_, err = resStmt.ExecContext(ctx, ins.Vendor, ins.Region, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return err
	}

	// 更新 host 表
	descStmt, err = tx.PrepareContext(ctx, UpdateDescribeSQL)
	if err != nil {
		return err
	}
	defer descStmt.Close()
	_, err = descStmt.ExecContext(ctx, ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
