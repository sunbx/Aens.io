package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ApiLoginController struct {
	BaseController
}

type ApiRegisterController struct {
	BaseController
}

type ApiNamesAuctionsController struct {
	BaseController
}

type ApiNamesPriceController struct {
	BaseController
}

type ApiNamesOverController struct {
	BaseController
}
type ApiNamesMyRegisterController struct {
	BaseController
}
type ApiNamesMyOverController struct {
	BaseController
}

type ApiNamesUpdateController struct {
	BaseController
}

type ApiNamesInfoController struct {
	BaseController
}
type ApiNamesAddController struct {
	BaseController
}
type ApiTransferAddController struct {
	BaseController
}
type ApiUserInfoController struct {
	BaseController
}

type Login struct {
	Code int64     `json:"code"`
	Data LoginData `json:"data"`
	Msg  string    `json:"msg"`
	Time int64     `json:"time"`
}

type LoginData struct {
	Address     string `json:"address"`
	Mnemonic    string `json:"mnemonic"`
	RedirectURI string `json:"redirectUri"`
	SigningKey  string `json:"signingKey"`
}

func (c *ApiLoginController) Post() {
	mnemonic := c.GetString("mnemonic")
	mnemonic = strings.Trim(mnemonic, " ")
	mnemonic = strings.Replace(mnemonic, "\n", "", -1)
	index := c.GetString("index")
	resp, err := http.PostForm("https://aeasy.io/api/user/login",
		url.Values{
			"app_id":        {beego.AppConfig.String("AEASY::appId")},
			"mnemonic":      {mnemonic},
			"index_address": {index},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	fmt.Println(string(body))
	var login Login
	err = json.Unmarshal([]byte(string(body)), &login)

	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return
	}
	c.Ctx.SetCookie("address", login.Data.Address)
	c.Ctx.WriteString(string(body))
}

func (c *ApiRegisterController) Post() {
	resp, err := http.PostForm("https://aeasy.io/api/user/register",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesAuctionsController) Post() {
	page := c.GetString("page")
	resp, err := http.PostForm("https://aeasy.io/api/names/auctions",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesPriceController) Post() {
	page := c.GetString("page")
	resp, err := http.PostForm("https://aeasy.io/api/names/price",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesOverController) Post() {
	page := c.GetString("page")
	resp, err := http.PostForm("https://aeasy.io/api/names/over",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesMyRegisterController) Post() {
	page := c.GetString("page")
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm("https://aeasy.io/api/names/my/register",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"page":       {page},
			"signingKey": {signingKey},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesMyOverController) Post() {
	page := c.GetString("page")
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm("https://aeasy.io/api/names/my/over",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"page":       {page},
			"signingKey": {signingKey},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesUpdateController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm("https://aeasy.io/api/names/update",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"signingKey": {signingKey},
			"name":       {name},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesInfoController) Post() {
	name := c.GetString("name")
	resp, err := http.PostForm("https://aeasy.io/api/names/info",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"name":   {name},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesAddController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm("https://aeasy.io/api/names/add",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"name":       {name},
			"signingKey": {signingKey},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		c.ErrorJson(-500, e.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiTransferAddController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	recipientAddress := c.GetString("recipientAddress")
	nameTransfer, done := c.TransferName(name, signingKey, recipientAddress)
	if done {
		return
	}
	c.SuccessJson(nameTransfer.Data.Hash)
}

func (c *ApiUserInfoController) Post() {
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm("https://aeasy.io/api/user/info",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"signingKey": {signingKey},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	c.Ctx.WriteString(string(body))
}
