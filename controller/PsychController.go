package controller

import (
	"fmt"

	"goweb/core"
)

type PsychController struct {
	core.App
}

// IndexAction ，第一个字母I，必须大写，这样的话，才是public访问权限
// 否则的话，在反射的时候，是获取不到这个方法的
func (i *PsychController) IndexAction() {
	resW := i.App.Context().RseW
	fmt.Fprintln(resW, "psychcontroller")
}

