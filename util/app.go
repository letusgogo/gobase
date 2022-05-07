package util

import (
	"fmt"
	"github.com/letusgogo/gobase/conf"
	"github.com/letusgogo/gobase/log"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/cmd"
	"go.uber.org/zap/zapcore"
	// 统一配置中心
	consulConfig "github.com/micro/go-plugins/config/source/consul/v2"

	// 加载服务发现,注册中心. 可以 backup 防止一个不行了
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
)

var (
	GAppName  = ""
	GAppEnv   = ""
	GLogLevel = zapcore.DebugLevel
)

func GetLogLevel() zapcore.Level {
	return GLogLevel
}

func SetLogLevel(levelStr string) {
	err := GLogLevel.UnmarshalText([]byte(levelStr))
	if err != nil {
		panic(err)
	}
}

func GetAppName() string {
	if GAppName == "" {
		panic("--app_name is empty")
	}
	return GAppName
}

func SetAppName(name string) {
	GAppName = name
}

func GetAppEnv() string {
	if GAppEnv == "" {
		panic("--app_env is empty")
	}
	return GAppEnv
}

func SetAppEnv(appEnv string) {
	GAppEnv = appEnv
}

func InitApp(myCmd cmd.Cmd) {
	InitConf()
	InitCmd(myCmd)
}

func InitLog() {
	log.InitLogWithPath("logs/"+GetAppEnv()+"/", GetAppName(), GetLogLevel())
}

func InitConf() {
	// 1. 加载配置文件
	consulSource := consulConfig.NewSource(
		// optionally specify consul address; default to localhost:8500
		consulConfig.WithAddress("consul.middleware.com:8500"),
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
	// 默认注册中心添加 consul, k8s 作为注册中心
	// 增加命令行
	updateCmd.App().Flags = append(cmd.DefaultFlags,
		&cli.StringFlag{
			Name:    "app_name",
			Usage:   "set app name",
			EnvVars: []string{"APP_NAME"},
		},
		&cli.StringFlag{
			Name:    "app_env",
			Usage:   "set app run app_environment",
			EnvVars: []string{"APP_ENV"},
		},
		&cli.StringFlag{
			Name:    "log_level",
			Usage:   "log level",
			EnvVars: []string{"LOG_LEVEL"},
		},
	)

	// 获取原来的 getRegistryFromConf 函数
	cmdBefore := cmd.App().Before
	cmd.App().Before = func(ctx *cli.Context) error {
		// 获取命令行中应用的基本参数
		getAppInfo(ctx)
		// 获取注册中心信息
		getRegistryFromConf(ctx)
		//初始化日志
		InitLog()
		// 还调用原函数
		if cmdBefore != nil {
			return cmdBefore(ctx)
		} else {
			return nil
		}
	}
	// 获取原来的 After 函数
	cmdAfter := cmd.App().After
	cmd.App().After = func(ctx *cli.Context) error {
		// 插入自定义的部分
		// 还调用原函数
		if cmdAfter != nil {
			return cmdAfter(ctx)
		} else {
			return nil
		}
	}
}

func getRegistryFromConf(ctx *cli.Context) {
	var err error

	// 如果命令行指定了则使用命令行指定的否则使用配置中心的 注册中心
	regTypeCmd := ctx.String("registry")
	regAddrCmd := ctx.String("registry_address")
	if regTypeCmd != "" {
		_ = ctx.Set("registry", regTypeCmd)
		_ = ctx.Set("registry_address", regAddrCmd)
	} else {
		// 从配置中心 middleware 获取 registry 配置
		registryConf := conf.RegistryConf{}
		err = config.Get(GetAppEnv(), conf.MiddlewareNamespace(), "registry").Scan(&registryConf)
		if err != nil {
			panic(err)
		}
		_ = ctx.Set("registry", registryConf.Type)
		_ = ctx.Set("registry_address", registryConf.Addr)
	}

	// 如果命令行指定了则使用命令行指定的否则使用配置中心的 broker
	brokerTypeCmd := ctx.String("broker")
	brokerAddressCmd := ctx.String("broker_address")
	if brokerTypeCmd != "" {
		_ = ctx.Set("broker", brokerTypeCmd)
		_ = ctx.Set("broker_address", brokerAddressCmd)
	} else {
		// 从配置中心 middleware 获取 broker 配置
		brokerConf := conf.BrokerConf{}
		err = config.Get(GetAppEnv(), conf.MiddlewareNamespace(), "broker").Scan(&brokerConf)
		if err != nil {
			panic(err)
		}
		if brokerConf.Type != "" {
			_ = ctx.Set("broker", brokerConf.Type)
			_ = ctx.Set("broker_address", brokerConf.Addr)
		}
	}

}

func getAppInfo(ctx *cli.Context) {
	// 命令行中获取基本信息
	appEnv := ctx.String("app_env")
	if appEnv == "" {
		panic("app_env is empty")
	}
	SetAppEnv(appEnv)
	fmt.Println("app_env:", appEnv)

	appName := ctx.String("app_name")
	if appName == "" {
		panic("app_name is empty")
	}
	SetAppName(appName)
	fmt.Println("app_name:", appName)

	levelStr := ctx.String("log_level")
	if levelStr == "" {
		panic("log_level is empty")
	}
	SetLogLevel(levelStr)
	fmt.Println("log_level:", levelStr)
}
