package model

import (
	"goweb/basedb"
	"goweb/tool/FyLogger"
)

// FrjjFamilyUser 凡人家教家风家训对应数据表
type FrjjFamilyUser struct {
	ID         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Openid     string    `db:"openid" json:"openid" xorm:"unique VARCHAR(100)"`
//	Surname    string    `xorm:"VARCHAR(10)"`
//	Name       string    `xorm:"VARCHAR(10)"`
//	Phone      string    `xorm:"VARCHAR(20)"`
	Info       string    `json:"info" xorm:"not null default '' VARCHAR(150)"`
//	Font       int       `xorm:"default 0 TINYINT(2)"`
//	Area       int       `xorm:"default 0 TINYINT(2)"`
//	Workplace  string    `xorm:"default '' VARCHAR(50)"`
//	ImgURL     string    `xorm:"VARCHAR(300)"`
//	FromOpenid string    `xorm:"VARCHAR(100)"`
//	IsSend     int       `xorm:"default 0 TINYINT(2)"`
	CreateTime string		`db:"create_time" json:"create_time" xorm:"DATETIME"`
	SendTime   string		`db:"send_time" json:"send_time" xorm:"DATETIME"`
	UpdateTime string 		`db:"update_time" json:"update_time" xorm:"DATETIME"`
}

// GetUser data
func (user *FrjjFamilyUser) GetUser(id int) (err error) {
	// 用了scan函数，会自动close连接
	err = basedb.WxDb.QueryRowx("SELECT id , info , create_time FROM frjj_family_user where id = ? limit 1", id).StructScan(user)
	if err != nil {
		FyLogger.DebugLog(err)
		return
	}
	return
}

func (u *FrjjFamilyUser) GetUserList() (user []FrjjFamilyUser) {
	// 数据多的时候，使用Select方法需要注意，因为会占用比较大的内存
	// 可使用Query循环
	basedb.WxDb.Select(&user, "SELECT id , info , create_time FROM frjj_family_user")
	//log.Println(user)
	return
}
