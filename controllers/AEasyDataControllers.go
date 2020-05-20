package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *HomeController) getNameBase() (NameBase, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/base",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return NameBase{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return NameBase{}, true
	}
	var nameBase NameBase
	err = json.Unmarshal([]byte(string(body)), &nameBase)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return NameBase{}, true
	}
	return nameBase, false
}

func (c *DetailController) getNameInfo(name string) (NameInfo, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/info",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"name":   {name},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return NameInfo{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return NameInfo{}, true
	}
	var nameInfo NameInfo
	err = json.Unmarshal([]byte(string(body)), &nameInfo)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return NameInfo{}, true
	}
	return nameInfo, false
}

type NameBase struct {
	Code int64        `json:"code"`
	Data NameBaseData `json:"data"`
	Msg  string       `json:"msg"`
	Time int64        `json:"time"`
}

type NameBaseData struct {
	Count    int64                 `json:"count"`
	Ranking  []NameBaseDataRanking `json:"ranking"`
	Sum      int64                 `json:"sum"`
	SumPrice float64               `json:"sum_price"`
}

type NameBaseDataRanking struct {
	Id       int     `json:"id"`
	NameNum  int64   `json:"name_num"`
	Owner    string  `json:"owner"`
	SumPrice float64 `json:"sum_price"`
	Address  string  `json:"address"`
}

//===================================================================

func (c *AuctionController) getAuctionName(page string) (Name, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/auctions",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Name{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var name Name
	err = json.Unmarshal([]byte(string(body)), &name)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return Name{}, true
	}
	return name, false
}

func (c *AuctionMyController) getAuctionMyName(page string, address string) (Name, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/my/over",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"page":    {page},
			"address": {address},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Name{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var name Name
	err = json.Unmarshal([]byte(string(body)), &name)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return Name{}, true
	}
	return name, false
}

func (c *ExpireMyController) getRegisterMyName(page string, address string) (Name, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/my/register",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"page":    {page},
			"address": {address},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Name{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var name Name
	err = json.Unmarshal([]byte(string(body)), &name)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return Name{}, true
	}
	return name, false
}

func (c *PriceController) getPriceName(page string) (Name, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/price",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Name{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var name Name
	err = json.Unmarshal([]byte(string(body)), &name)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return Name{}, true
	}
	return name, false
}

func (c *ExpireController) getExpireName(page string) (Name, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/over",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Name{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var name Name
	err = json.Unmarshal([]byte(string(body)), &name)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return Name{}, true
	}
	return name, false
}

func (c *BaseController) getUserInfo(address string) (UserInfo, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/user/info",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"address": {address},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return UserInfo{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var userInfo UserInfo
	err = json.Unmarshal([]byte(string(body)), &userInfo)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return UserInfo{}, true
	}
	return userInfo, false
}

type UserInfo struct {
	Code int64        `json:"code"`
	Data UserInfoData `json:"data"`
	Msg  string       `json:"msg"`
	Time int64        `json:"time"`
}

type UserInfoData struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type Name struct {
	Code int64      `json:"code"`
	Data []NameData `json:"data"`
	Msg  string     `json:"msg"`
	Time int64      `json:"time"`
}

type NameData struct {
	Id            int     `json:"id"`
	CurrentHeight int64   `json:"current_height"`
	Price         string  `json:"price"`
	CurrentPrice  string  `json:"current_price"`
	EndHeight     int64   `json:"end_height"`
	Length        int64   `json:"length"`
	Name          string  `json:"name"`
	OverHeight    int64   `json:"over_height"`
	Owner         string  `json:"owner"`
	Address       string  `json:"address"`
	StartHeight   int64   `json:"start_height"`
	ThHash        string  `json:"th_hash"`
	Progress      int64   `json:"progress"`
	Countdown     string  `json:"countdown"`
	Gains         float64 `json:"gains"`
}

//========================

type NameInfo struct {
	Code int64        `json:"code"`
	Data NameInfoData `json:"data"`
	Msg  string       `json:"msg"`
	Time int64        `json:"time"`
}

type NameInfoDataClaim struct {
	AccountID string `json:"account_id"`
	Fee       int64  `json:"fee"`
	Time      int64  `json:"time"`

	Name        string  `json:"name"`
	NameFee     string  `json:"name_fee"`
	NameSalt    float64 `json:"name_salt"`
	Nonce       int64   `json:"nonce"`
	Type        string  `json:"type"`
	Version     int64   `json:"version"`
	Gains       float64 `json:"gains"`
	Address     string  `json:"address"`
	Date        string  `json:"date"`
	BlockHeight int64   `json:"block_height"`
}

type NameInfoData struct {
	Claim         []NameInfoDataClaim `json:"claim"`
	CurrentHeight int64               `json:"current_height"`
	CurrentPrice  string              `json:"current_price"`
	Price         string              `json:"price"`
	EndHeight     int64               `json:"end_height"`
	Length        int64               `json:"length"`
	Name          string              `json:"name"`
	OverHeight    int64               `json:"over_height"`
	Owner         string              `json:"owner"`
	StartHeight   int64               `json:"start_height"`
	ThHash        string              `json:"th_hash"`
}
