package util

var (
	appName = ""
	appEnv  = ""
)

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
