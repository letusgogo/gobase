package util

import "fmt"

const middlewareNamespace = "middleware"

type MidMysqlConf struct {
	Host         string `json:"Host"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxOpenConns int    `json:"MaxOpenConns"`
}

type MidRedisConf struct {
	Addr        string `json:"Addr"`
	Password    string `json:"Password"`
	PoolSize    int    `json:"PoolSize"`
	MaxRetries  int    `json:"MaxRetries"`
	IdleTimeout int64  `json:"IdleTimeout"`
}

type MidInfluxDBConf struct {
	Addr string `json:"Addr"`
}

type MidKafkaConf struct {
	Addr string `json:"Addr"`
}

func GetMiddlewareNamespace() string {
	return middlewareNamespace
}

func GetMysqlUrl(host, user, pass, dBName string) string {
	dBOption := "charset=utf8&parseTime=True&loc=Local"
	result := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", user, pass, host, dBName, dBOption)
	return result
}
