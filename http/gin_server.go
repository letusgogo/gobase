package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

// GinServer 启动一个httpserver 对外提供服务。会依赖各个组件的业务系统
type GinServer struct {
	port       uint16
	ginEngine  *gin.Engine
	httpServer *http.Server
}

func NewGinServer(port uint16) *GinServer {
	ginEngine := gin.Default()

	return &GinServer{
		port:      port,
		ginEngine: ginEngine,
		httpServer: &http.Server{
			Handler: ginEngine,
		},
	}
}

func (h *GinServer) GinGroup(relativePath string) *gin.RouterGroup {
	return h.ginEngine.Group(relativePath)
}

func (h *GinServer) GinEngine() *gin.Engine {
	return h.ginEngine
}

// Start 会阻塞
func (h *GinServer) Start() error {
	// 设置服务器监听请求端口
	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", h.port))
	if err != nil {
		return err
	}

	err = h.httpServer.Serve(l)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}

func (h *GinServer) Stop() error {
	withTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	err := h.httpServer.Shutdown(withTimeout)
	if err != nil {
		return err
	} else {
		return nil
	}
}
