package tool

import (
	"blogV2.zozoo.net/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

//注册xorm
var Orm *xorm.Engine

func OrmEngine() (*xorm.Engine, error) {

	//获取配置文件
	config := Conf.Database

	//注册数据库
	engine, err := xorm.NewEngine(
		config.Driver,
		config.User+":"+config.Password+"@("+config.Host+")/"+config.DbName+"?charset="+config.Charset)
	if err != nil {
		return nil, err
	}

	//判断执行sql操作时是否显示sql语句
	engine.ShowSQL(config.ShowSql)

	//映射数据表
	err = engine.Sync2(
		new(model.User),
		new(model.Article),
		new(model.Tag),
		new(model.TagArticle),
		)

	if err != nil {
		return nil, err
	}

	Orm = engine
	return engine, nil
}
