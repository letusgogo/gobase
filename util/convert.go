package util

import (
	"errors"
	"reflect"
	"strconv"
)

func GetBoolean(val interface{}) (bool, error) {
	switch val.(type) {
	case bool:
		return val.(bool), nil
	case int, uint, int8, uint8, int32, uint32, int64, uint64, float32, float64:
		{
			getInt64, err := GetInt64(val)
			if err != nil {
				return false, err
			}
			if getInt64 == 0 {
				return false, nil
			} else if getInt64 == 1 {
				return true, nil
			} else {
				return false, errors.New("not 0 or 1 can not convert boolean")
			}
		}
	case string:
		if val.(string) == "true" {
			return true, nil
		} else if val.(string) == "false" {
			return false, nil
		} else if val.(string) == "1" {
			return true, nil
		} else if val.(string) == "0" {
			return false, nil
		} else {
			return false, errors.New("can not convert string to boolean")
		}
	default:
		return false, errors.New("value type:" + typeof(val) + " can not convert to int")
	}
}

func GetBooleanWithDefault(val interface{}, defaultVal bool) bool {
	boolVal, err := GetBoolean(val)
	if err != nil {
		return defaultVal
	}

	return boolVal
}

func GetInt8(val interface{}) (int8, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}

	return int8(i), nil
}

func GetInt8WithDefault(val interface{}, defaultVal int8) int8 {
	getInt8, err := GetInt8(val)
	if err != nil {
		return defaultVal
	} else {
		return getInt8
	}
}

func GetUint8(val interface{}) (uint8, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}

	if i < 0 {
		return 0, errors.New("negative can not convert unsigned number")
	}

	return uint8(i), nil
}

func GetUint8WithDefault(val interface{}, defaultVal uint8) uint8 {
	getUint8, err := GetUint8(val)
	if err != nil {
		return defaultVal
	} else {
		return getUint8
	}
}

func GetInt16(val interface{}) (int16, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}

	return int16(i), nil
}

func GetInt16WithDefault(val interface{}, defaultVal int16) int16 {
	getInt16, err := GetInt16(val)
	if err != nil {
		return defaultVal
	} else {
		return getInt16
	}
}

func GetUint16(val interface{}) (uint16, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}

	if i < 0 {
		return 0, errors.New("negative can not convert unsigned number")
	}

	return uint16(i), nil
}

func GetUint16WithDefault(val interface{}, defaultVal uint16) uint16 {
	getUint16, err := GetUint16(val)
	if err != nil {
		return defaultVal
	} else {
		return getUint16
	}
}

func GetInt32(val interface{}) (int32, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}

	return int32(i), nil
}

func GetInt32WithDefault(val interface{}, defaultVal int32) int32 {
	getInt32, err := GetInt32(val)
	if err != nil {
		return defaultVal
	} else {
		return getInt32
	}
}

func GetUInt32(val interface{}) (uint32, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}

	if i < 0 {
		return 0, errors.New("negative can not convert unsigned number")
	}

	return uint32(i), nil
}

func GetUint32WithDefault(val interface{}, defaultVal uint32) uint32 {
	uInt32, err := GetUInt32(val)
	if err != nil {
		return defaultVal
	} else {
		return uInt32
	}
}

func GetUint64(val interface{}) (uint64, error) {
	// 把值统一转为 int
	i, err := GetInt64(val)
	if err != nil {
		return 0, err
	}
	if i < 0 {
		return 0, errors.New("negative can not convert unsigned number")
	}

	return uint64(i), nil
}

func GetUint64WithDefault(val interface{}, defaultVal uint64) uint64 {
	getUint64, err := GetUint64(val)
	if err != nil {
		return defaultVal
	} else {
		return getUint64
	}
}

func GetInt64WithDefault(val interface{}, defaultVal int64) int64 {
	getInt64, err := GetInt64(val)
	if err != nil {
		return defaultVal
	} else {
		return getInt64
	}
}

func GetInt64(val interface{}) (int64, error) {
	if val == nil {
		return 0, errors.New("nil pointer")
	}
	switch val.(type) {
	case int:
		return int64(val.(int)), nil
	case uint:
		return int64(val.(uint)), nil
	case int8:
		return int64(val.(int8)), nil
	case uint8:
		return int64(val.(uint8)), nil
	case int16:
		return int64(val.(int16)), nil
	case uint16:
		return int64(val.(uint16)), nil
	case int32:
		return int64(val.(int32)), nil
	case uint32:
		return int64(val.(uint32)), nil
	case int64:
		return val.(int64), nil
	case uint64:
		// 不考虑越界问题
		return int64(val.(uint64)), nil
	case string:
		str, _ := val.(string)
		if intVal, e := strconv.Atoi(str); e != nil {
			return 0, e
		} else {
			return int64(intVal), nil
		}
	case float64:
		// float 差越界判断
		float, _ := val.(float64)
		if isInteger(float) {
			return int64(float), nil
		} else {
			return 0, errors.New(strconv.FormatFloat(float, 'e', 5, 64) + " is not integer")
		}
	case float32:
		float, _ := val.(float32)
		if isInteger(float64(float)) {
			return int64(float), nil
		} else {
			return 0, errors.New(strconv.FormatFloat(float64(float), 'e', 5, 32) + " is not integer")
		}
	default:
		return 0, errors.New("value type:" + typeof(val) + " can not convert to int")
	}
}

func GetStringWithDefault(val interface{}, defaultVal string) string {
	getString, err := GetString(val)
	if err != nil {
		return defaultVal
	} else {
		return getString
	}
}

func GetString(val interface{}) (string, error) {
	if val == nil {
		return "", errors.New("nil pointer")
	}
	switch val.(type) {
	case string:
		return val.(string), nil
	case bool:
		valBool, _ := val.(bool)
		return strconv.FormatBool(valBool), nil
	case int8:
		valInt, _ := val.(int8)
		return strconv.FormatInt(int64(valInt), 10), nil
	case uint8:
		valInt, _ := val.(uint8)
		return strconv.FormatInt(int64(valInt), 10), nil
	case int16:
		valInt, _ := val.(int16)
		return strconv.FormatInt(int64(valInt), 10), nil
	case uint16:
		valInt, _ := val.(uint16)
		return strconv.FormatInt(int64(valInt), 10), nil
	case int32:
		valInt, _ := val.(int32)
		return strconv.FormatInt(int64(valInt), 10), nil
	case uint32:
		valInt, _ := val.(uint32)
		return strconv.FormatInt(int64(valInt), 10), nil
	case int64:
		valInt, _ := val.(int64)
		return strconv.FormatInt(valInt, 10), nil
	case uint64:
		valInt, _ := val.(uint64)
		return strconv.FormatInt(int64(valInt), 10), nil
	case float32:
		valFloat, _ := val.(float32)
		return strconv.FormatFloat(float64(valFloat), 'e', 5, 32), nil
	case float64:
		valFloat, _ := val.(float64)
		return strconv.FormatFloat(valFloat, 'e', 5, 64), nil
	default:
		return "", errors.New("value type:" + typeof(val) + " can not convert to string dp")
	}
}

// 获取类型名
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// 判断是否是整数 float
func isInteger(a float64) bool {
	// float64(int64(a)) 会截取小数点之前的部分
	if a-float64(int64(a)) == 0 {
		return true
	} else {
		return false
	}
}
