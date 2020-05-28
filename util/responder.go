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

// 从定义好的 ErrInfo 构造错误对象。可以替换掉里面的 Msg
// errInfo := CopyErrInfo(ErrDataBase)
// errInfo.Msg = "redis error"
func CopyErrInfo(err *ErrInfo) *ErrInfo {
	return &ErrInfo{Ret: err.Ret, Msg: err.Msg, Error: err.Error}
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

func WriteRpcRsp(rspPtr interface{}, err RpcError, datas map[string]interface{}) {
	if nil == datas {
		datas = make(map[string]interface{})
	}
	datas["Ret"] = err.GetRet()
	datas["Msg"] = err.GetMsg()
	SetStructVals(rspPtr, datas)
}
