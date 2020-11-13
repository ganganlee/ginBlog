package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//配置文件获取中心
type (
	Config struct {
		AppName  string `json:"app_name"`
		AppModel string `json:"app_model"`
		AppHost  string `json:"app_host"`
		AppPort  string `json:"app_port"`
		Database Database
	}

	//数据库配置中心
	Database struct {
		Driver   string `json:"driver"`
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DbName   string `json:"db_name"`
		Charset  string `json:"charset"`
		ShowSql  bool   `json:"show_sql"`
	}
)

var Conf *Config

func ParseConfig(path string) *Config {
	conf := new(Config)

	//读取配置文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(path)
		panic("打开配置文件出错")
	}
	defer file.Close()

	confByte, err := ioutil.ReadAll(file)
	if err != nil {
		panic("读取配置文件出错")
	}

	if err := json.Unmarshal(confByte, conf); err != nil {
		panic("解析配置文件出错")
	}

	Conf = conf
	return conf
}
