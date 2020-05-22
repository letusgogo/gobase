package util

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/cmd"
	consulConfig "github.com/micro/go-plugins/config/source/consul"
	consultRegistry "github.com/micro/go-plugins/registry/consul"
)

var (
	appName = ""
	appEnv  = ""
)

type Registry struct {
	Type    string
	Address string
}

func GetAppName() string {
	if appName == "" {
		panic("appName is empty")
	}
	return appName
}

func SetAppName(name string) {
	appName = name
}

func GetEnv() string {
	if appEnv == "" {
		panic("appEnv is empty")
	}
	return appEnv
}

func SetEnv(env string) {
	appEnv = env
}

func InitApp(myCmd cmd.Cmd) {
	InitConf()
	InitCmd(myCmd)
}

func InitConf() {
	// 1. 加载配置文件
	consulSource := consulConfig.NewSource(
		// optionally specify consul address; default to localhost:8500
		consulConfig.WithAddress("middleware.consul.com:8500"),
		// optionally specify prefix; defaults to /micro/config
		consulConfig.WithPrefix("config/"),
		// optionally strip the provided prefix from the keys, defaults to false
		consulConfig.StripPrefix(true),
	)
	// 2. 从指定的源加载
	if err := config.Load(consulSource); err != nil {
		panic(err)
		return
	}
}

func InitCmd(updateCmd cmd.Cmd) {
	// 默认注册中心添加 consul 作为注册中心
	cmd.DefaultRegistries["consul"] = consultRegistry.NewRegistry
	// 增加命令行
	updateCmd.App().Flags = append(cmd.DefaultFlags,
		cli.StringFlag{
			Name:   "app_name",
			Usage:  "set app name",
			EnvVar: "APP_NAME",
		},
		cli.StringFlag{
			Name:   "env",
			Usage:  "set app run environment",
			EnvVar: "ENV",
		},
	)
	// 替换从命令行输入参数的方式创建 registry
	cmdBefore := cmd.App().Before
	cmd.App().Before = func(ctx *cli.Context) error {
		registry := Registry{}
		// 从配置中心获取注册中心信息
		err := config.Get(GetEnv(), GetAppName(), "registry").Scan(&registry)
		err = ctx.Set("registry", registry.Type)
		err = ctx.Set("registry_address", registry.Address)
		if err != nil {
			panic(err)
		}
		env := ctx.String("env")
		fmt.Println("env:", env)
		SetEnv(env)

		appName := ctx.String("app_name")
		fmt.Println("app_name:", appName)
		SetAppName(appName)

		// 还调用原函数
		return cmdBefore(ctx)
	}
}
