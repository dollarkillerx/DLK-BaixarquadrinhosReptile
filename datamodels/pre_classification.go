package datamodels

type PreClassification struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Name       string `xorm:"not null default '' VARCHAR(266)" form:"name"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') VARCHAR(255)" form:"language"`
	Describe   string `xorm:"not null default '' comment('漫画描述') VARCHAR(366)" form:"describe"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
}
