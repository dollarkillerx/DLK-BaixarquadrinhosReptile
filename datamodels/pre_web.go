package datamodels

type PreWeb struct {
	Id               int                  `xorm:"not null pk autoincr INT(10)"`
	Language         string               `xorm:"not null default '' comment('语言编号缩写') unique VARCHAR(255)" form:"language"`
	HotCategories    string               `xorm:"not null comment('热门分类') VARCHAR(355)" form:"hot_categories"`
	TheBest          string               `xorm:"not null comment('最好') VARCHAR(355)" form:"the_best"`
	MostConcerned    string               `xorm:"not null comment('最受关注') VARCHAR(355)" form:"most_concerned"`
	RecentlyUpdated  string               `xorm:"not null comment('最近更新') VARCHAR(355)" form:"recently_updated"`
	Search           string               `xorm:"not null comment('搜索') VARCHAR(255)" form:"search"`
	Copyright        string               `xorm:"not null comment('版权注释') VARCHAR(255)" form:"copyright"`
	WatchNow         string               `xorm:"not null comment('立即观看') VARCHAR(355)" form:"watch_now"`
	SignIn           string               `xorm:"not null comment('登录') VARCHAR(355)" form:"sign_in"`
	Comic            string               `xorm:"not null comment('漫画') VARCHAR(355)" form:"comic"`
	WebName          string               `xorm:"not null comment('网站名称') VARCHAR(355)" form:"web_name"`
	WebLogo          string               `xorm:"not null comment('网站logo') VARCHAR(355)" form:"web_logo"`
	NextPage         string               `xorm:"not null comment('下一页') VARCHAR(355)" form:"previous_page"`
	PreviousPage     string               `xorm:"not null comment('上一页') VARCHAR(355)" form:"next_page"`
	PageTurning      string               `xorm:"not null comment('翻页') VARCHAR(366)" form:"page_turning"`
	Serializing      string               `xorm:"not null default '' comment('连载中') VARCHAR(255)" form:"serializing"`
	End              string               `xorm:"not null default '' comment('完结') VARCHAR(255)" form:"end"`
	UserLogin        string               `xorm:"not null default '' comment('会员登录') VARCHAR(255)" form:"user_login"`
	Phone            string               `xorm:"not null default '' comment('请输入电话号码') VARCHAR(355)" form:"phone"`
	VerificationCode string               `xorm:"not null default '' comment('请输入验证码') VARCHAR(355)" form:"verification_code"`
	BackToContents   string               `xorm:"not null comment('返回目录') VARCHAR(255)" form:"BackToContents"`
	BackToHome       string               `xorm:"not null comment('返回首页') VARCHAR(255)" form:"BackToHome"`
	Return           string               `xorm:"not null comment('返回') VARCHAR(255)" form:"Return"`
	Logout           string               `xorm:"not null comment('登出') VARCHAR(255)" form:"Logout"`
	Categories       string               `xorm:"not null comment('分类') VARCHAR(255)" form:"Categories"`
	Nav              []*PreNav            `xorm:"-"` // 导航
	Cls              []*PreClassification `xorm:"-"` // 分类
	S3Url            string               `xorm:"-"` // s3 url地址
}
