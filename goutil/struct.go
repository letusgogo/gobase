package goutil

import "reflect"

func SetStructVal(rsp interface{}, datas map[string]interface{}) {
	if nil == datas {
		return
	}
	rType := reflect.TypeOf(rsp)
	// rsp 要为指针才能改变值
	if rType.Kind() != reflect.Ptr {
		panic("not a ptr")
	}
	// 获取指针指向的原始的值
	rVal := reflect.ValueOf(rsp).Elem()
	// 获取指针指向元素类型
	rType = rType.Elem()
	if rType.Kind() != reflect.Struct {
		panic("not struct ptr")
	}

	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if data, ok := datas[t.Name]; ok {
			dV := reflect.ValueOf(data)
			if dV.Type().Name() != f.Type().Name() {
				panic("type is not same," + "arrt:" + t.Name + "type is " + t.Type.Name() + ", data type is " + dV.Type().Name())
			} else {
				f.Set(dV)
			}
		}
	}
}
