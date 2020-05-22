package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"time"
	"github.com/beego/i18n"
)

type BaseController struct {
	i18n.Locale
	beego.Controller
}

func (this *BaseController) Prepare() {
	var languageCookie = this.Ctx.GetCookie("language")
	this.Lang = "en-US"
	if languageCookie == "zh-CN" {
		this.Lang = "zh-CN"
	} else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

type ReturnMsg struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Stime int64       `json:"time"`
	Data  interface{} `json:"data"`
}

type JsonData struct {
}

var bm, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

func (c *BaseController) SuccessJson(data interface{}) {
	serviceTime := time.Now().UnixNano() / 1e6
	res := ReturnMsg{
		200, "success", serviceTime, data,
	}
	jsons, _ := json.Marshal(res)

	c.Ctx.WriteString(string(jsons))
}

func (c *BaseController) ErrorJson(code int, msg string, data interface{}) {
	serviceTime := time.Now().UnixNano() / 1e6
	res := ReturnMsg{
		code, msg, serviceTime, data,
	}
	jsons, _ := json.Marshal(res)

	c.Ctx.WriteString(string(jsons))
}

