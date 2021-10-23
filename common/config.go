package common

import (
	"fmt"
	"os"
)

type App struct {
	Port        string
	JwtSecret   string
	ErrLogPath  string
	infoLogPath string
}

type Mysql struct {
	Username string
	Password string
	Host     string
	Dbname   string
	Port     string
}

type Redis struct {
	Host     string
	Prot     string
	Username string
}

type config struct {
	App   App
	Mysql Mysql
	Redis Redis
}

var Config config

func ConfigInit() {
	// 读取当前路径下的conf
	GOENV := os.Getenv("GO_ENV")
	fmt.Println(GOENV)

	Config = config{
		App: App{
			Port:        ":8180",
			JwtSecret:   "22126a93-3f84-4a5d-a4ce-a9b8ddde6bb2",
			ErrLogPath:  "",
			infoLogPath: "",
		},
		Mysql: Mysql{
			Username: "root",
			Password: "",
			Host:     "127.0.0.1",
			Dbname:   "tanzi1997",
			Port:     "3306",
		},
	}
}
