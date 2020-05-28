package http

import (
	"bytes"
	"git.iothinking.com/base/gobase/log"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// http 服务器
func TestNewHttpServer(t *testing.T) {
	log.InitLog("httpserver", zapcore.DebugLevel)
	hand := HandlerMap{
		"/a": handA,
		"/b": handB,
	}
	server := NewHttpServer(":8081", true, hand)
	go server.StartHttpServer(5, 5)
	server.WaitForTerminal()
}

//func TestNewMultiHttpServer(t *testing.T) {
//	hand1 := HandlerMap{
//		"/a": handA,
//		"/b": handB,
//	}
//	server1 := NewHttpServer(":8081", hand1)
//	go server1.StartHttpServer()
//	go server1.WaitForTerminal()
//
//	hand2 := HandlerMap{
//		"/c": handC,
//		"/d": handD,
//	}
//
//	server2 := NewHttpServer(":8082", hand2)
//	go server2.StartHttpServer()
//	server2.WaitForTerminal()
//}

// 下载文件
//func TestNewHttpFileServer(t *testing.T) {
//	hand := HandlerMap{
//		"/file":handFile,
//	}
//	server := NewHttpServer(":8081", hand)
//	go server.StartHttpServer()
//	server.WaitForTerminal()
//}

func handA(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("A"))
}

func handB(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("B"))
}

func handC(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("C"))
}

func handD(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("D"))
}

// 下载一个图片
func handFile(w http.ResponseWriter, req *http.Request) {
	//
	fileBytes, err := ioutil.ReadFile("README.md")
	if err != nil {
		_, _ = w.Write([]byte("open file failed"))
		return
	}
	sendFile(w, req, fileBytes)
}

// 返回图片
func sendFile(w http.ResponseWriter, req *http.Request, picBytes []byte) {
	// 构造 http 头
	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Pragma", "public")
	// 把图片内容放到 io.ReadSeeker 里。
	picReader := bytes.NewReader(picBytes)

	// 发送数据
	utc := time.Now().UTC()
	http.ServeContent(
		w,
		req,
		"db_motion_1555461224.jpeg",
		utc,
		picReader)
}
