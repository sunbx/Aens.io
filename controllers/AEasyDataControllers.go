package controllers

import (
	"encoding/json"
	"fmt"
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

func (c *BaseController) getNameInfo(name string) (NameInfo, bool) {
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
	if nameInfo.Code != 200 {
		c.ErrorJson(-100, nameInfo.Msg, JsonData{})
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

func (c *BaseController) TransferName(name string, signingKey string, recipientAddress string) (NameTransfer, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/transfer",
		url.Values{
			"app_id":           {beego.AppConfig.String("AEASY::appId")},
			"name":             {name},
			"signingKey":       {signingKey},
			"recipientAddress": {recipientAddress},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return NameTransfer{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return NameTransfer{}, true
	}
	var nameTransfer NameTransfer
	err = json.Unmarshal([]byte(string(body)), &nameTransfer)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return NameTransfer{}, true
	}
	if nameTransfer.Code != 200 {
		c.ErrorJson(-100, nameTransfer.Msg, JsonData{})
		return NameTransfer{}, true
	}
	return nameTransfer, false
}

func TransferName(name string, signingKey string, recipientAddress string) (NameTransfer, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/names/transfer",
		url.Values{
			"app_id":           {beego.AppConfig.String("AEASY::appId")},
			"name":             {name},
			"signingKey":       {signingKey},
			"recipientAddress": {recipientAddress},
		})
	if err != nil {
		return NameTransfer{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NameTransfer{}, true
	}
	var nameTransfer NameTransfer
	err = json.Unmarshal([]byte(string(body)), &nameTransfer)
	if err != nil {
		return NameTransfer{}, true
	}
	if nameTransfer.Code != 200 {
		return NameTransfer{}, true
	}
	return nameTransfer, false
}

func (c *BaseController) getTh(th string) (Th, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/ae/th_hash",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"th":     {th},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Th{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return Th{}, true
	}
	var thm Th
	err = json.Unmarshal([]byte(string(body)), &thm)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return Th{}, true
	}
	if thm.Code != 200 {
		c.ErrorJson(-100, thm.Msg, JsonData{})
		return Th{}, true
	}
	return thm, false
}

func GetTh(th string) (Th, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/ae/th_hash",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"th":     {th},
		})
	if err != nil {
		return Th{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Th{}, true
	}
	var thm Th
	err = json.Unmarshal([]byte(string(body)), &thm)
	if err != nil {
		return Th{}, true
	}
	if thm.Code != 200 {
		return Th{}, true
	}
	return thm, false
}

func Transfer(amount string, address string, signingKey string) (WalletTransfer, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/wallet/transfer",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"amount":     {amount},
			"address":    {address},
			"signingKey": {signingKey},
		})
	if err != nil {
		fmt.Println("transfer error", err)
		return WalletTransfer{}, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("transfer error", err)
		return WalletTransfer{}, true
	}
	var walletTransfer WalletTransfer
	err = json.Unmarshal([]byte(string(body)), &walletTransfer)
	if err != nil {
		fmt.Println("transfer error", err)
		return WalletTransfer{}, true
	}
	if walletTransfer.Code != 200 {
		fmt.Println("transfer error", walletTransfer)
		return WalletTransfer{}, true
	}
	return walletTransfer, false
}

func (c *BaseController) getBlockHeight() (int64, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/ae/block_top",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return -1, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return -1, true
	}
	var blockHeight BlockHeight
	err = json.Unmarshal([]byte(string(body)), &blockHeight)
	if err != nil {
		c.ErrorJson(-100, err.Error(), JsonData{})
		return -1, true
	}
	if blockHeight.Code != 200 {
		c.ErrorJson(-100, blockHeight.Msg, JsonData{})
		return -1, true
	}
	return blockHeight.Data.Height, false
}

func GetBlockHeight() (int64, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/ae/block_top",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		return -1, true
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, true
	}
	var blockHeight BlockHeight
	err = json.Unmarshal([]byte(string(body)), &blockHeight)
	if err != nil {
		return -1, true
	}
	if blockHeight.Code != 200 {
		return -1, true
	}
	return blockHeight.Data.Height, false
}

func (c *BaseController) getRegisterMyName(page string, address string) (Name, bool) {
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

func (c *BaseController) getUserInfoAddress(address string) (UserInfo, bool) {
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

func (c *BaseController) getUserInfoSigningKey(signingKey string) (UserInfo, bool) {
	resp, err := http.PostForm("https://aeasy.io/api/user/info",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"signingKey": {signingKey},
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
	if userInfo.Code != 200 {
		c.ErrorJson(-100, userInfo.Msg, JsonData{})
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

//=====================================

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

//==================================
type NameTransfer struct {
	Code int64            `json:"code"`
	Data NameTransferData `json:"data"`
	Msg  string           `json:"msg"`
	Time int64            `json:"time"`
}

type NameTransferData struct {
	BlockHash   string             `json:"BlockHash"`
	BlockHeight int64              `json:"BlockHeight"`
	Error       interface{}        `json:"Error"`
	Hash        string             `json:"Hash"`
	Mined       bool               `json:"Mined"`
	Signature   string             `json:"Signature"`
	SignedTx    string             `json:"SignedTx"`
	Tx          NameTransferDataTx `json:"Tx"`
}

type NameTransferDataTx struct {
	AccountID    string `json:"AccountID"`
	AccountNonce int64  `json:"AccountNonce"`
	Fee          int64  `json:"Fee"`
	NameID       string `json:"NameID"`
	RecipientID  string `json:"RecipientID"`
	TTL          int64  `json:"TTL"`
}

//==============================================
type Th struct {
	Code int64  `json:"code"`
	Data ThData `json:"data"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

type ThDataTx struct {
	AccountID   string `json:"account_id"`
	Fee         int64  `json:"fee"`
	NameID      string `json:"name_id"`
	Nonce       int64  `json:"nonce"`
	RecipientID string `json:"recipient_id"`
	TTL         int64  `json:"ttl"`
	Type        string `json:"type"`
	Version     int64  `json:"version"`
}

type ThData struct {
	BlockHash   string   `json:"block_hash"`
	BlockHeight int64    `json:"block_height"`
	Hash        string   `json:"hash"`
	Signatures  []string `json:"signatures"`
	Tx          ThDataTx `json:"tx"`
}

//========================================
type BlockHeight struct {
	Code int64           `json:"code"`
	Data BlockHeightData `json:"data"`
	Msg  string          `json:"msg"`
	Time int64           `json:"time"`
}

type BlockHeightData struct {
	Height int64 `json:"height"`
}

//=====================================
type WalletTransfer struct {
	Code int64              `json:"code"`
	Data WalletTransferData `json:"data"`
	Msg  string             `json:"msg"`
	Time int64              `json:"time"`
}

type WalletTransferDataTxTx struct {
	Amount      float64 `json:"Amount"`
	Fee         int64   `json:"Fee"`
	Nonce       int64   `json:"Nonce"`
	Payload     string  `json:"Payload"`
	RecipientID string  `json:"RecipientID"`
	SenderID    string  `json:"SenderID"`
	TTL         int64   `json:"TTL"`
}

type WalletTransferDataTx struct {
	BlockHash   string                 `json:"BlockHash"`
	BlockHeight int64                  `json:"BlockHeight"`
	Error       interface{}            `json:"Error"`
	Hash        string                 `json:"Hash"`
	Mined       bool                   `json:"Mined"`
	Signature   string                 `json:"Signature"`
	SignedTx    string                 `json:"SignedTx"`
	Tx          WalletTransferDataTxTx `json:"Tx"`
}

type WalletTransferData struct {
	Tx WalletTransferDataTx `json:"tx"`
}
