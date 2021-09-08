package util

import (
	"fmt"
	"git.iothinking.com/base/gobase/conf"
	"git.iothinking.com/base/gobase/log"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/cmd"
	consulConfig "github.com/micro/go-plugins/config/source/consul/v2"
	consultRegistry "github.com/micro/go-plugins/registry/consul/v2"
	"go.uber.org/zap/zapcore"
)

var (
	appName  = ""
	appEnv   = ""
	logLevel = zapcore.DebugLevel
)

func GetLogLevel() zapcore.Level {
	return logLevel
}

func SetLogLevel(levelStr string) {
	err := logLevel.UnmarshalText([]byte(levelStr))
	if err != nil {
		panic(err)
	}
}

func GetAppName() string {
	if appName == "" {
		panic("--app_name is empty")
	}
	return appName
}

func SetAppName(name string) {
	appName = name
}

func GetAppEnv() string {
	if appEnv == "" {
		panic("--app_env is empty")
	}
	return appEnv
}

func SetAppEnv(app_env string) {
	appEnv = app_env
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
	// 默认注册中心添加 consul 作为注册中心
	cmd.DefaultRegistries["consul"] = consultRegistry.NewRegistry
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
	appEnv := ctx.String("appEnv")
	if appEnv == "" {
		panic("appEnv is empty")
	}
	SetAppEnv(appEnv)
	fmt.Println("appEnv:", appEnv)

	appName := ctx.String("app_name")
	if appName == "" {
		panic("appName is empty")
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
