package protocol

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/middleware/cors"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lewinloo/restful-api-demo/apps"
	"github.com/lewinloo/restful-api-demo/conf"
)

func NewHttpService() *HttpService {
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1Mb
		Addr:              conf.C().App.HttpAddr(),
		Handler:           cors.AllowAll().Handler(r),
	}

	return &HttpService{
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

type HttpService struct {
	r      gin.IRouter
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

func (s *HttpService) Start() error {
	// 加载 Handler，把所有模块注册给 Gin
	// 注册IOC的http handler
	apps.InitGin(s.r)

	// 已加载app的日志信息
	apps := apps.LoadedGinApps()
	s.l.Infof("loaded gin apps: %v", apps)

	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service stopped success")
			return nil
		}
		return err
	}
	return nil
}

func (s *HttpService) Stop() {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("shutdown http service error, %s", err)
	}
}
