package goconf

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewConfigLoaderForApollo(t *testing.T) {
	loader, _ := NewConfigLoader("apollo")
	properties, err := loader.Load("http://10.0.200.5:8181/configs/httpdns/default/APP_ALL")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(properties)
	}

	result := properties.GetList("UsRegularRules", nil)
	fmt.Println(result)
}

type regionInfoStu map[string](map[string](map[string]string))

func TestNewConfigLoaderForConfig(t *testing.T) {

	var regionInfo regionInfoStu

	//noinspection GoRedundantParens
	loader, _ := NewConfigLoader("apollo")
	properties, err := loader.Load("http://172.16.248.36:8181/configs/httpdns/region-info/APP_ALL")
	if err != nil {
		t.Fatal(err)
	}

	bytes := properties.GetString("region", "")

	fmt.Println(string(bytes))

	err = json.Unmarshal([]byte(bytes), &regionInfo)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProperties_getString_FromApollo(t *testing.T) {
	// 从apollo
	loader, _ := NewConfigLoader(FromApollo)
	properties, _ := loader.Load("http://10.0.200.5:8181/configs/httpdns/default/APP_ALL")

	val := properties.GetString("RedisPass", "")

	fmt.Println(val)
}

func TestProperties_getString_FromJsonFile(t *testing.T) {
	loader, _ := NewConfigLoader(FromJsonFile)
	properties, _ := loader.Load("./config.json")

	val := properties.GetString("RedisPass", "")

	fmt.Println(val)
}

func TestProperties_getInt_FromApollo(t *testing.T) {
	// 从apollo
	loader, _ := NewConfigLoader(FromApollo)
	properties, _ := loader.Load("http://172.16.248.36:8181/configs/httpdns/default/APP_ALL")

	val := properties.GetInt("PoolSize", 1000)

	fmt.Println(val)
}

func TestProperties_getInt_FromsonFile(t *testing.T) {
	// 从文件
	loader, _ := NewConfigLoader(FromJsonFile)
	properties, _ := loader.Load("./config.json")

	val := properties.GetInt("PoolCap", 1000)

	fmt.Println(val)
}

func TestProperties_GetBool_FromApollo(t *testing.T) {
	// 从apollo
	loader, _ := NewConfigLoader(FromApollo)
	properties, _ := loader.Load("http://172.16.248.36:8181/configs/encrypt-pic/default/APP_ALL")
	fmt.Printf("conf:%v\n", properties.RawData())
	val := properties.GetBool("Debug", false)

	fmt.Println(val)
}

func TestProperties_GetBool_FromJsonFile(t *testing.T) {
	// 从文件
	loader, _ := NewConfigLoader(FromJsonFile)
	properties, _ := loader.Load("./config.json")

	val := properties.GetBool("Debug", false)

	fmt.Println(val)
}

func Test_jsonFile_LoadToStruct(t *testing.T) {
	type ConfInfo struct {
		ListenAddr string `json:"ListenAddr"`
		PoolSize   int    `json:"PoolSize"`
		RedisAddr  string `json:"RedisAddr"`
		RedisDB    int    `json:"RedisDB"`
		RedisPass  string `json:"RedisPass"`
	}

	confInfo := ConfInfo{}

	loader, _ := NewConfigLoader(FromJsonFile)
	err := loader.LoadToStruct("./test_config.json", &confInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(confInfo)
}

func Test_apollo_LoadToStruct(t *testing.T) {
	type ApolloInfo struct {
		AppID          string `json:"appId"`
		Cluster        string `json:"cluster"`
		Configurations struct {
			ListenAddr string `json:"ListenAddr"`
			PoolSize   string `json:"PoolSize"`
			RedisAddr  string `json:"RedisAddr"`
			RedisDB    string `json:"RedisDB"`
			RedisPass  string `json:"RedisPass"`
		} `json:"configurations"`
		NamespaceName string `json:"namespaceName"`
		ReleaseKey    string `json:"releaseKey"`
	}

	apolloInfo := ApolloInfo{}

	loader, _ := NewConfigLoader(FromApollo)
	err := loader.LoadToStruct("http://172.16.248.36:8181/configs/httpdns/default/APP_ALL", &apolloInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(apolloInfo)
}
