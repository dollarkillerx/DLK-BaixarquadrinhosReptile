package datamodels

type PreUser struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Number     string `xorm:"not null default '' comment('phone') unique VARCHAR(255)" form:"number"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') VARCHAR(255)"`
	Password   string `xorm:"not null default '' CHAR(32)" form:"password"`
	Salt       string `xorm:"not null default '' CHAR(32)"`
	Name       string `xorm:"not null default '' VARCHAR(255)"`
	Admin      int    `xorm:"not null default 0 comment('0普通用户1管理员') TINYINT(4)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
}
