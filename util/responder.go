package util

// rcp 响应接口
type RpcResponder interface {
	SetRet(ret int32)
	SetMsg(msg string)
}

type RpcError interface {
	GetRet() int32
	GetMsg() string
}

// 定义错误
type ErrInfo struct {
	Ret   int32  // 错误码
	Msg   string // 展示给用户看的
	Error error  // 保存内部错误信息
}

func (e *ErrInfo) GetRet() int32 {
	return e.Ret
}

func (e *ErrInfo) GetMsg() string {
	return e.Msg
}

//ret=0 成功。
var (
	ErrNot    = &ErrInfo{0, "success", nil}
	ErrUnknow = &ErrInfo{1, "unknow error", nil}

	ErrParam    = &ErrInfo{2, "param error", nil}
	ErrDataBase = &ErrInfo{3, "database error", nil}
)

func WriteRpcRsp(rspPtr interface{}, rpcError RpcError, data map[string]interface{}) {
	if nil == data {
		data = make(map[string]interface{})
	}
	if rpcError == nil {
		panic("rpcError is nil")
	}
	data["Ret"] = rpcError.GetRet()
	data["Msg"] = rpcError.GetMsg()
	SetStructVals(rspPtr, data)
}

func WriteRpcRspWithMsg(rspPtr interface{}, rpcError RpcError, msg string, data map[string]interface{}) {
	if nil == data {
		data = make(map[string]interface{})
	}
	if rpcError == nil {
		panic("rpcError is nil")
	}
	data["Ret"] = rpcError.GetRet()
	data["Msg"] = msg
	SetStructVals(rspPtr, data)
}
