package goutil

import (
	"fmt"
	"testing"
)

// 设备的 dp 点
type DpInfo struct {
	DpId    int32  `protobuf:"varint,2,opt,name=dpId,proto3" json:"dpId"`
	Ename   string `protobuf:"bytes,3,opt,name=ename,proto3" json:"ename"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Type    string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Mode    string `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode"`
	Value   string `protobuf:"bytes,7,opt,name=value,proto3" json:"value"`
	RawData []byte `protobuf:"bytes,8,opt,name=rawData,proto3" json:"rawData"`
}

// 响应数据点请求
type GetDpRsp struct {
	Ret   int32     `protobuf:"varint,1,opt,name=ret,proto3" json:"ret"`
	Msg   string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	DevId string    `protobuf:"bytes,3,opt,name=devId,proto3" json:"devId"`
	Dps   []*DpInfo `protobuf:"bytes,4,rep,name=dps,proto3" json:"dps"`
}

func (g *GetDpRsp) SetRet(ret int32) {
	g.Ret = ret
}

func (g *GetDpRsp) SetMsg(msg string) {
	g.Msg = msg
}

func TestGeneratorRpcRsp(t *testing.T) {
	rsp := new(GetDpRsp)
	dps := make([]*DpInfo, 0)
	dp1 := DpInfo{
		DpId:    0,
		Ename:   "yuhaiyang",
		Name:    "于",
		RawData: nil,
	}

	dp2 := DpInfo{
		DpId:    1,
		Ename:   "pear",
		Name:    "周",
		RawData: nil,
	}
	dps = append(dps, &dp1)
	dps = append(dps, &dp2)

	GeneratorRpcRsp(rsp, ErrNot, map[string]interface{}{
		"DevId": "d12312",
		"Dps":   dps,
	})

	fmt.Println(rsp)
	if rsp.Ret != 0 {
		t.Error("rsp.Ret is not 0")
	}

	if rsp.Msg != "success" {
		t.Error("rsp.Ret is not success")
	}

	if len(rsp.Dps) != 2 {
		t.Error("rsp.Dps is not equal 2")
	}
}
