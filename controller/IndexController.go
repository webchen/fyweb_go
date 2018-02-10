package controller

import (
	"strconv"

	"goweb/core"
	"goweb/model"
	"goweb/tool"
)

// IndexController , the default
type IndexController struct {
	core.App
}

/*
var indexCtl = new(*IndexController)
indexCtl.Before(indexCtl.Context())
*/

// IndexAction ，第一个字母I，必须大写，这样的话，才是public访问权限
// 否则的话，在反射的时候，是获取不到这个方法的
func (i *IndexController) IndexAction() {
	i.Context().Echo(tool.GetIp(i.Context().Req))
}

func (i *IndexController) DataAction() {
	id, _ := tool.GetInt(i.App.Context().Req, "id")
	user := new(model.FrjjFamilyUser)

	err := user.GetUser(id)
	if err != nil {
		i.Context().SetDataCode(1102).SetDataMessage("id为 " + strconv.Itoa(id) + " 的数据不存在").Show()
		return
	}
	i.Context().SetDataCode(1101).SetDataMessage("获取数据成功").SetData(user).Show()
}

func (i *IndexController) ListAction() {
	user := new(model.FrjjFamilyUser)
	list := user.GetUserList()
	i.Context().SetDataCode(1001).SetDataMessage("").SetData(list).Show()
}
