package util

var (
	appName = "micro-app"
	appEnv  = "dev"
)

func GetAppName() string {
	return appName
}

func SetAppName(name string) {
	appName = name
}

func GetEnv() string {
	return appEnv
}

func SetEnv(env string) {
	appEnv = env
}
