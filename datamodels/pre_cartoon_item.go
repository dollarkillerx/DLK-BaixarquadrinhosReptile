package datamodels

import "github.com/MobileCPX/PreComic4You/defs"

type PreCartoonItem struct {
	Id         int                `xorm:"not null pk autoincr INT(10)"`
	Name       string             `xorm:"not null default '' VARCHAR(366)" form:"name"`
	Collection int                `xorm:"not null default 0 comment('集') INT(10)" form:"collection"`
	CartoonId  int                `xorm:"not null default 0 comment('漫画id') INT(10)" form:"cartoon_id"`
	Content    string             `xorm:"comment('图片主体,JSON') TEXT" form:"content"`
	Read       int                `xorm:"not null default 0 comment('阅读量') INT(10)" form:"read"`
	CreateTime int                `xorm:"not null default 0 comment('创建时间') INT(10)" form:"create_time"`
	UpdateTime int                `xorm:"not null default 0 comment('更新时间') INT(10)" form:"update_time"`
	DeleteTime int                `xorm:"not null default 0 comment('软删除') index INT(10)" form:"delete_time"`
	Language   string             `xorm:"not null comment('语言 简称') VARCHAR(255)" form:"language"`
	Item       *defs.CartoonItems `xorm:"-"`
}
