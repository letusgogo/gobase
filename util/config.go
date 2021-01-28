package util

import "fmt"

const middlewareNamespace = "middleware"

type MysqlConf struct {
	Host         string `json:"Host"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxOpenConns int    `json:"MaxOpenConns"`
}

type RedisConf struct {
	Addr        string `json:"Addr"`
	Password    string `json:"Password"`
	PoolSize    int    `json:"PoolSize"`
	MaxRetries  int    `json:"MaxRetries"`
	IdleTimeout int64  `json:"IdleTimeout"`
}

// influxdb 配置
type InfluxDBConf struct {
	Addr string `json:"Addr"`
}

// kafka 配置
type KafkaConf struct {
	Addr string `json:"Addr"`
}

// 短信配置
type SmsConf struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	SignName        string `json:"SignName"`
	TemplateCode    string `json:"TemplateCode"`
}

// 服务的 Broker 配置
type BrokerConf struct {
	Type string
	Addr string
}

// mqtt 配置
type MqttConf struct {
	Brokers []string
}

// 注册中心相关的配置
type RegistryConf struct {
	Type string `json:"Type"`
	Addr string `json:"Addr"`
}

func MiddlewareNamespace() string {
	return middlewareNamespace
}

func GetMysqlUrl(host, user, pass, dBName string) string {
	dBOption := "charset=utf8&parseTime=True&loc=Local"
	result := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", user, pass, host, dBName, dBOption)
	return result
}
