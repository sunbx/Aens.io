package main

import (
	"github.com/beego/i18n"
	_ "names/routers"
	"github.com/astaxie/beego"
)

func main() {
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}

