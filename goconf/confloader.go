package goconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

// 配置类型
const (
	FromJsonFile = "jsonFile"
	FromApollo   = "apollo"
)

// 加载配置文件的接口
type ConfLoader interface {
	// 到目标地址拉配置,解析为 Properties
	Load(uri string) (Properties, error)
	// 解析到 struct
	LoadToStruct(uri string, dstStruct interface{}) error
}

// 配置文件解析
type Properties struct {
	rawData map[string]interface{}
}

func (p *Properties) RawData() map[string]interface{} {
	return p.rawData
}

func (p *Properties) SetRawData(rawData map[string]interface{}) {
	p.rawData = rawData
}

func (p *Properties) GetBytes(name string, defaultVal []byte) []byte {
	data, ok := p.rawData[name]
	if !ok {
		return defaultVal
	}

	bytes, err := json.Marshal(&data)
	if err != nil {
		return defaultVal
	}

	return bytes
}

func (p *Properties) GetList(name string, defaultVal []interface{}) []interface{} {
	str := p.GetString(name, "")
	if len(str) == 0 {
		return defaultVal
	}

	result := make([]interface{}, 0)
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return defaultVal
	}

	return result
}

func (p *Properties) GetString(name, defaultVal string) string {
	val, ok := p.rawData[name]
	if !ok {
		return defaultVal
	}

	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	case int64:
		return strconv.FormatInt(val.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'E', -1, 32)
	case bool:
		if val.(bool) {
			return "true"
		} else {
			return "false"
		}
	default:
		panic("不支持的类型转 string," + reflect.TypeOf(val).String())
	}
}

func (p *Properties) GetBool(name string, defaultVal bool) bool {
	val, ok := p.rawData[name]
	if !ok {
		return defaultVal
	}

	switch val.(type) {
	case bool:
		return val.(bool)
	case int:
		return val.(int) != 0
	case string:
		i, err := strconv.ParseBool(val.(string))
		if err != nil {
			return defaultVal
		} else {
			return i
		}
	default:
		panic(errors.New("除了bool, string,int 暂不支持其他类型转换到 bool." + reflect.TypeOf(val).String()))
	}
}

// 如果是浮点数可能会有精度损失
func (p *Properties) GetInt(name string, defaultVal int) int {
	val, ok := p.rawData[name]
	if !ok {
		return defaultVal
	}

	switch val.(type) {
	case int:
		return val.(int)
	case float64:
		return int(val.(float64))
	case string:
		i, err := strconv.Atoi(val.(string))
		if err != nil {
			return defaultVal
		} else {
			return i
		}
	default:
		panic(errors.New("除了int,string,float64 暂不支持其他类型转换到 int." + reflect.TypeOf(val).String()))
	}
}

// 如果是浮点数可能会有精度损失
func (p *Properties) GetInt64(name string, defaultVal int64) int64 {
	val, ok := p.rawData[name]
	if !ok {
		return defaultVal
	}

	switch val.(type) {
	case int64:
		return val.(int64)
	case float64:
		return int64(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return defaultVal
		} else {
			return i
		}
	default:
		panic(errors.New("除了int64,string,float64 暂不支持其他类型转换到  int." + reflect.TypeOf(val).String()))
	}
}

func NewConfigLoader(logType string) (ConfLoader, error) {
	switch logType {
	case FromJsonFile:
		return new(jsonFile), nil

	case FromApollo:
		return new(apollo), nil

	default:

	}
	return nil, errors.New("未识别的配置类型")
}

////////////////////////// 从配置文件加载配置 ////////////////////////
type jsonFile struct {
}

func (*jsonFile) Load(uri string) (Properties, error) {
	data, err := ioutil.ReadFile(uri)
	if err != nil {
		return Properties{}, err
	}

	rawMap := make(map[string]interface{})
	if err = json.Unmarshal(data, &rawMap); err != nil {
		return Properties{}, err
	}

	return Properties{rawData: rawMap}, nil
}

func (*jsonFile) LoadToStruct(uri string, dstStruct interface{}) error {
	data, err := ioutil.ReadFile(uri)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dstStruct)
}

//////////////////从apollo加载配置文件//////////////////
type apollo struct {
}

func (*apollo) Load(uri string) (Properties, error) {
	// http 请求 apollo
	resp, err := http.Get(uri)
	if err != nil {
		return Properties{}, err
	}

	// 从 http 获取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Properties{}, err
	}

	// 解析为 map
	tmpMap := make(map[string]interface{})
	if err = json.Unmarshal(body, &tmpMap); err != nil {
		return Properties{}, err
	}

	//从apollo 未读取到配置
	if status, ok := tmpMap["status"]; ok {
		return Properties{}, fmt.Errorf("not found status: %v", status)
	}

	// 获取真正的配置数据
	configureData := tmpMap["configurations"]
	confDataMap, ok := configureData.(map[string]interface{})
	if !ok {
		return Properties{}, errors.New("configData can not convert string")
	}

	return Properties{confDataMap}, nil
}

func (*apollo) LoadToStruct(uri string, dstStruct interface{}) error {
	// http 请求 apollo
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	// 从 http 获取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 解析为 map
	return json.Unmarshal(body, dstStruct)
}
