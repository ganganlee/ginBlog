package main
//个人博客 v2

import (
	"blogV2.zozoo.net/route"
	"blogV2.zozoo.net/tool"
	"github.com/gin-gonic/gin"
	"os"
	"sync"
)

func main() {

	//读取配置文件
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tool.ParseConfig(dir+"/config/config.json")

	//实例化gin框架
	r := gin.Default()

	//实例化数据库
	_,err = tool.OrmEngine()
	if err != nil {
		panic(err)
	}

	//注册路由
	route.InitRoute(r)

	//注册携程等待
	tool.Wg = &sync.WaitGroup{}

	r.Run(":9920") // listen and serve on 0.0.0.0:8080
}