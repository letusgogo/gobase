package goutil

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
	ErrUnknow = &ErrInfo{100, "unknow error", nil}

	ErrParam = &ErrInfo{1, "param error", nil}
)

func WriteRpcRsp(rsp interface{}, err RpcError, datas map[string]interface{}) {
	datas["Ret"] = err.GetRet()
	datas["Msg"] = err.GetMsg()
	SetStructVal(rsp, datas)
}
