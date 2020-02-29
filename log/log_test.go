package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

type GroupInfo struct {
	GroupId string  `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId"`
	Ename   string  `protobuf:"bytes,2,opt,name=ename,proto3" json:"ename"`
	Name    string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Desc    string  `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc"`
	DpIds   []int32 `protobuf:"varint,5,rep,packed,name=dpIds,proto3" json:"dpIds"`
}

type UpdateGroupByDevReq struct {
	DevId  string       `protobuf:"bytes,1,opt,name=devId,proto3" json:"devId"`
	Groups []*GroupInfo `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups"`
}

func TestInitLog(t *testing.T) {
	dpIds := make([]int32, 0)
	dpIds = append(dpIds, 1, 2, 3)
	info := &GroupInfo{
		GroupId: "uuid",
		Ename:   "en",
		Name:    "中文",
		Desc:    "desc",
		DpIds:   dpIds,
	}

	groups := make([]*GroupInfo, 0)
	groups = append(groups, info)
	req := UpdateGroupByDevReq{
		DevId:  "dev",
		Groups: groups,
	}
	InitLog("test", zapcore.DebugLevel)
	log.Info("test", zap.Any("req", &req))
}
