# 读取配置

## 本地 json 文件

```
    loader, _ := NewConfigLoader("jsonFile")
	properties, err := loader.Load("../conf/config.json")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(properties)
	}
	
```

## 远程 apollo


```
func TestNewConfigLoaderForApollo(t *testing.T) {
	loader, _ := NewConfigLoader("apollo")
	properties, err := loader.Load("http://10.0.200.5:8181/configs/httpdns/default/APP_ALL")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(properties)
	}

}

```

## 转为 struct
```
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
	err := loader.LoadToStruct("http://10.0.200.5:8181/configs/httpdns/default/APP_ALL", &apolloInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(apolloInfo)
}

```

