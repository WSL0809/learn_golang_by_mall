package conf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode string
	AppPort string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
)

func init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(fmt.Sprintf("Fail to read file :%v", err))
	}
	LoadServer(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	AppPort = file.Section("server").Key("AppPort").String()

	
}