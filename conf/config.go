package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const middlewareNamespace = "middleware"

type MysqlConf struct {
	Host         string `json:"Host"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxOpenConns int    `json:"MaxOpenConns"`
}

func (m *MysqlConf) String() string {
	b, err := json.Marshal(*m)
	if err != nil {
		return fmt.Sprintf("%+v", *m)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *m)
	}
	return out.String()
}

type RedisConf struct {
	Addr        string `json:"Addr"`
	Password    string `json:"Password"`
	PoolSize    int    `json:"PoolSize"`
	MaxRetries  int    `json:"MaxRetries"`
	IdleTimeout int64  `json:"IdleTimeout"`
}

func (r *RedisConf) String() string {
	b, err := json.Marshal(*r)
	if err != nil {
		return fmt.Sprintf("%+v", *r)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *r)
	}
	return out.String()
}

// influxdb 配置
type InfluxDBConf struct {
	Addr string `json:"Addr"`
}

func (i *InfluxDBConf) String() string {
	b, err := json.Marshal(*i)
	if err != nil {
		return fmt.Sprintf("%+v", *i)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *i)
	}
	return out.String()
}

// kafka 配置
type KafkaConf struct {
	Addr string `json:"Addr"`
}

func (k *KafkaConf) String() string {
	b, err := json.Marshal(*k)
	if err != nil {
		return fmt.Sprintf("%+v", *k)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *k)
	}
	return out.String()
}

// 短信配置
type SmsConf struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	SignName        string `json:"SignName"`
	TemplateCode    string `json:"TemplateCode"`
	Debug           bool   `json:"Debug"`
}

func (s *SmsConf) String() string {
	b, err := json.Marshal(*s)
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	return out.String()
}

// 配置
type TelConf struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	TemplateCode    string `json:"TemplateCode"`
	Debug           bool   `json:"Debug"`
}

func (t *TelConf) String() string {
	b, err := json.Marshal(*t)
	if err != nil {
		return fmt.Sprintf("%+v", *t)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *t)
	}
	return out.String()
}

// 服务的 Broker 配置
type BrokerConf struct {
	Type string
	Addr string
}

func (c *BrokerConf) String() string {
	b, err := json.Marshal(*c)
	if err != nil {
		return fmt.Sprintf("%+v", *c)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *c)
	}
	return out.String()
}

// mqtt 配置
type MqttConf struct {
	Brokers []string
}

func (c *MqttConf) String() string {
	b, err := json.Marshal(*c)
	if err != nil {
		return fmt.Sprintf("%+v", *c)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *c)
	}
	return out.String()
}

// 注册中心相关的配置
type RegistryConf struct {
	Type string `json:"Type"`
	Addr string `json:"Addr"`
}

func (c *RegistryConf) String() string {
	b, err := json.Marshal(*c)
	if err != nil {
		return fmt.Sprintf("%+v", *c)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *c)
	}
	return out.String()
}

type AppMysqlConf struct {
	User         string `json:"user" mapstructure:"host"`
	Pass         string `json:"pass" mapstructure:"pass"`
	Debug        bool   `json:"debug" mapstructure:"debug"`
	Host         string `json:"host" mapstructure:"host"`
	DBName       string `json:"db_name" mapstructure:"db_name"`
	MaxIdleConns int    `json:"maxIdle_conns"  mapstructure:"maxIdle_conns"`
	MaxOpenConns int    `json:"max_open_conns" mapstructure:"max_open_conns"`
}

func (m *AppMysqlConf) String() string {
	b, err := json.Marshal(*m)
	if err != nil {
		return fmt.Sprintf("%+v", *m)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *m)
	}
	return out.String()
}

func MiddlewareNamespace() string {
	return middlewareNamespace
}

func (m *AppMysqlConf) GetMysqlUrl() string {
	dBOption := "charset=utf8&parseTime=True&loc=Local"
	result := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", m.User, m.Pass, m.Host, m.DBName, dBOption)
	return result
}

// 图形验证码
type Captcha struct {
	KeyLong   int32 `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int32 `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int32 `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}
