package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type Msg struct {
	Code int
	Msg  string
}

func (e *Msg) String() string {
	if e == nil {
		return ""
	}
	b, err := json.Marshal(*e)
	if err != nil {
		return fmt.Sprintf("%+v", *e)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *e)
	}
	return out.String()
}

type Cmd struct {
}

func (c *Cmd) Start(context *gin.Context) {
	req := new(Msg)
	if err := context.BindJSON(req); err != nil {
		context.JSON(200, &Msg{
			Code: 1,
			Msg:  "msg not json",
		})
		return
	}
	fmt.Printf("recv:%s", req)
	context.JSON(200, &Msg{
		Code: 0,
		Msg:  "start",
	})
}

func (c *Cmd) Stop(context *gin.Context) {
	req := new(Msg)
	if err := context.BindJSON(req); err != nil {
		context.JSON(200, &Msg{
			Code: 1,
			Msg:  "msg not json",
		})
		return
	}
	fmt.Printf("recv:%s", req)
	context.JSON(200, &Msg{
		Code: 0,
		Msg:  "stop",
	})
}

type Stat struct {
}

func (c *Stat) Start(context *gin.Context) {
	req := new(Msg)
	if err := context.BindJSON(req); err != nil {
		context.JSON(200, &Msg{
			Code: 1,
			Msg:  "msg not json",
		})
		return
	}
	fmt.Printf("recv:%s", req)
	context.JSON(200, &Msg{
		Code: 0,
		Msg:  "start",
	})
}

func (c *Stat) Stop(context *gin.Context) {
	req := new(Msg)
	if err := context.BindJSON(req); err != nil {
		context.JSON(200, &Msg{
			Code: 1,
			Msg:  "msg not json",
		})
		return
	}
	fmt.Printf("recv:%s", req)
	context.JSON(200, &Msg{
		Code: 0,
		Msg:  "stop",
	})
}

func TestNewGinServer(t *testing.T) {
	ginServer := NewGinServer(8888)
	cmdGroup := ginServer.GinGroup("/cmd")

	// 相关接口
	{
		cmd := &Cmd{}
		cmdGroup.POST("/start", cmd.Start)
		cmdGroup.POST("/stop", cmd.Stop)
	}

	statGroup := ginServer.GinGroup("/manager")
	{
		stat := &Stat{}
		statGroup.POST("/start", stat.Start)
		statGroup.POST("/stop", stat.Stop)
	}

	// 启动服务
	go func(tin *testing.T) {
		err := ginServer.Start()
		if err != nil {
			tin.Error(err)
			return
		}
	}(t)

	time.Sleep(time.Second * 1)
	// 发送请求
	taskRequest := &Msg{
		Code: 0,
		Msg:  "hello",
	}
	sendData, _ := json.Marshal(taskRequest)
	hc := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := hc.Post(
		"http://127.0.0.1:8888/cmd/start",
		"application/json;charset=utf-8",
		bytes.NewBuffer(sendData),
	)

	if err != nil {
		t.Fatal(err)
	}

	//goland:noinspection ALL
	defer resp.Body.Close()
	// 解析响应
	body, _ := ioutil.ReadAll(resp.Body)
	taskResponse := new(Msg)
	err = json.Unmarshal(body, taskResponse)
	if err != nil {
		t.Fatal(err)
	}
	if taskResponse.Code != 0 {
		t.Error(taskResponse.Msg)
	}

	if err := ginServer.Stop(); err != nil {
		t.Fatal(err)
	}
}
