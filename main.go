package main

import (
	"github.com/iothink/gobase/log"
	"github.com/iothink/gobase/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 创建一个 user 请求
type CreateUserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Telphone             string   `protobuf:"bytes,2,opt,name=telphone,proto3" json:"telphone"`
	Pass                 string   `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass"`
	Pid                  int32    `protobuf:"varint,4,opt,name=pid,proto3" json:"pid"`
	Code                 string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

var ErrUserDuplicate = &util.ErrInfo{Ret: 100, Msg: "Duplicate user"}

func main() {
	req := &CreateUserReq{
		Name:     "于海洋",
		Telphone: "18989893637",
		Pass:     "123456",
		Pid:      10,
		Code:     "12312",
	}
	log.InitLog("test", zapcore.DebugLevel)
	log.Info("test", zap.Any("req", req), zap.Any("errr", ErrUserDuplicate))
}