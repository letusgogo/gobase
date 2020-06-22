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
//noinspection ALL
var (
	ErrNot    = &ErrInfo{0, "success", nil}
	ErrUnknow = &ErrInfo{1, "unknown error", nil}

	ErrParam    = &ErrInfo{2, "param error", nil}
	ErrDataBase = &ErrInfo{3, "database error", nil}
)

// 从 info 创建一个新的 ErrInfo 类型的对象。
// 当 msg 不为空,则用 msg 替换原 msg
// 当 err 不为 nil,则用 err 替换 原 err
//noinspection ALL
func NewErrInfo(info *ErrInfo, msg string, err error) *ErrInfo {
	errInfo := &ErrInfo{
		Ret:   info.Ret,
		Msg:   info.Msg,
		Error: info.Error,
	}
	if msg != "" {
		errInfo.Msg = msg
	}

	if err != nil {
		errInfo.Error = err
	}
	return errInfo
}

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

//noinspection ALL
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
