package gohttp

import (
	"context"
	"fmt"
	"github.com/iothink/gobase/golog"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// http 请求的回调
type HandlerMap map[string]func(w http.ResponseWriter, req *http.Request)

type SimpleHttpServer struct {
	accessLog  bool
	server     *http.Server
	addr       string
	handlerMap HandlerMap
}

func NewHttpServer(addr string, accessLog bool, handMap HandlerMap) *SimpleHttpServer {
	return &SimpleHttpServer{addr: addr, accessLog: accessLog, handlerMap: handMap}
}

// 开启 http 服务
func (hs *SimpleHttpServer) StartHttpServer(readTimeout, writeTimeout time.Duration) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hs.allHandler)

	hs.server = &http.Server{
		Addr:         hs.addr,
		WriteTimeout: time.Second * writeTimeout,
		ReadTimeout:  time.Second * readTimeout,
		Handler:      mux,
	}
	fmt.Printf("start httpserver port:(%s)\n", hs.addr)
	golog.Info("start httpserver ", zap.String("addr", hs.addr))
	err := hs.server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("httpserver error port:(%s),err:%s\n", hs.addr, err.Error())
			golog.Info("httpserver error ", zap.String("addr", hs.addr), zap.Error(err))
			panic(err)
		} else {
			fmt.Printf("exit httpserver port:(%s) ok,%s\n", hs.addr, err.Error())
			golog.Info("exit httpserver ", zap.String("addr", hs.addr), zap.Error(err))
		}
	}
}

// 关闭 http 服务
func (hs *SimpleHttpServer) WaitForTerminal() {
	// 接收系统信号
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)
	sig := <-osSig
	fmt.Println("recv:" + sig.String())
	//使用context控制srv.Shutdown的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := hs.server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

func (hs *SimpleHttpServer) allHandler(w http.ResponseWriter, req *http.Request) {
	// 如果 发现 panic，判断错误输出错误，否则 继续往上层 panic
	defer func() {
		if r := recover(); r != nil {
			golog.Error("发生未捕获错误:", zap.Any("r", r))
			http.Error(
				w, //向writer汇报错误
				http.StatusText(http.StatusInternalServerError), //错误描述信息（字符串）
				http.StatusInternalServerError)                  //系统内部错误
		}
	}()
	if hs.accessLog {
		golog.Debug("Access log",
			zap.Any("Head", req.Header),
			zap.Any("Path", req.URL.Path),
			zap.Any("RawQuery", req.URL.RawQuery))
	}

	path := req.URL.Path

	if _, ok := hs.handlerMap[path]; !ok {
		http.Error(w, "not find", http.StatusBadRequest)
		return
	}

	hs.handlerMap[path](w, req)
}
