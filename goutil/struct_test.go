package goutil

import (
	"fmt"
	"testing"
)

// 设备的 dp 点
type DpInfo2 struct {
	DpId    int32  `protobuf:"varint,2,opt,name=dpId,proto3" json:"dpId"`
	Ename   string `protobuf:"bytes,3,opt,name=ename,proto3" json:"ename"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Type    string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Mode    string `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode"`
	Value   string `protobuf:"bytes,7,opt,name=value,proto3" json:"value"`
	RawData []byte `protobuf:"bytes,8,opt,name=rawData,proto3" json:"rawData"`
}

// 响应数据点请求
type GetDpRsp2 struct {
	Ret   int32      `protobuf:"varint,1,opt,name=ret,proto3" json:"ret"`
	Msg   string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	DevId string     `protobuf:"bytes,3,opt,name=devId,proto3" json:"devId"`
	Dps   []*DpInfo2 `protobuf:"bytes,4,rep,name=dps,proto3" json:"dps"`
}

// 响应数据点请求
type GetDpRsp3 struct {
	Ret   int32      `protobuf:"varint,1,opt,name=ret,proto3" json:"ret"`
	Msg   string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	DevId string     `protobuf:"bytes,3,opt,name=devId,proto3" json:"devId"`
	Dps   []*DpInfo2 `protobuf:"bytes,4,rep,name=dps,proto3" json:"dps"`
}

func (g *GetDpRsp2) SetRet(ret int32) {
	g.Ret = ret
}

func (g *GetDpRsp2) SetMsg(msg string) {
	g.Msg = msg
}

func TestSetStructVal(t *testing.T) {
	rsp := new(GetDpRsp2)
	dps := make([]*DpInfo2, 0)
	dp1 := DpInfo2{
		DpId:    0,
		Ename:   "yuhaiyang",
		Name:    "于",
		RawData: nil,
	}

	dp2 := DpInfo2{
		DpId:    0,
		Ename:   "pear",
		Name:    "周",
		RawData: nil,
	}
	dps = append(dps, &dp1)
	dps = append(dps, &dp2)
	SetStructVals(rsp,
		map[string]interface{}{
			"Dps":   dps,
			"Ret":   int32(1),
			"Msg":   "succ",
			"DevId": "d12321"},
	)

	fmt.Println(*rsp)
}

func TestStructCopy(t *testing.T) {
	dstRsp := new(GetDpRsp2)
	srcRsp := new(GetDpRsp3)

	dps := make([]*DpInfo2, 0)
	dp1 := DpInfo2{
		DpId:    0,
		Ename:   "yuhaiyang",
		Name:    "于",
		RawData: nil,
	}

	dp2 := DpInfo2{
		DpId:    0,
		Ename:   "pear",
		Name:    "周",
		RawData: nil,
	}
	dps = append(dps, &dp1)
	dps = append(dps, &dp2)

	srcRsp.Ret = 0
	srcRsp.Msg = "succ"
	srcRsp.Dps = dps
	StructCopy(dstRsp, srcRsp, "Msg", "Dps")
	if dstRsp.Msg != "" {
		t.Error("Msg shoud empty")
	}
	if dstRsp.Dps != nil {
		t.Error("Dps shoud nil")
	}

	StructCopy(dstRsp, srcRsp)
	if dstRsp.Msg == "" {
		t.Error("Msg shoud not empty")
	}
	if dstRsp.Dps == nil {
		t.Error("Dps shoud not empty")
	}

	fmt.Println(*dstRsp)
}
