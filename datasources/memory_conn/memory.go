package memory_conn

import (
	"errors"
	"github.com/MobileCPX/PreComic4You/config"
	"github.com/MobileCPX/PreComic4You/datamodels"
	"github.com/MobileCPX/PreComic4You/datasources/mysql_conn"
	"sync"
)

var (
	InitWeb    *datamodels.PreWeb
	FractionDb *sync.Map
)

func init() {
	init_web()
	init_fraction_db()
}

// web 初始化
func init_web() {
	// 初始化语言
	InitWeb = &datamodels.PreWeb{}
	b, i := mysql_conn.MysqlEngine.Where("language = ?", config.MyConfig.App.Language).Get(InitWeb)
	if i != nil || b == false {
		panic(errors.New("缺少语言包"))
	}
}

// 初始化分数数据库
func init_fraction_db() {
	FractionDb = &sync.Map{}
}
