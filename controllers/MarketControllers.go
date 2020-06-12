package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"names/models"
	"names/utils"
	"strconv"
	"sync"
)

type MarketInController struct {
	BaseController
}

type MarketCancelController struct {
	BaseController
}

type MarketOutController struct {
	BaseController
}

var lock sync.Mutex

func (c *MarketOutController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")

	nameInfo, done := c.getNameInfo(name)
	//域名有问题或者其他错误
	if done {
		c.ErrorJson(-500, "names error", JsonData{})
		return
	}
	//如果域名还没到账户
	if nameInfo.Data.CurrentHeight <= nameInfo.Data.EndHeight {
		c.ErrorJson(-500, "The domain name is still up for grabs", JsonData{})
		return
	}

	//如果域名已经过期
	if nameInfo.Data.CurrentHeight >= nameInfo.Data.OverHeight {
		c.ErrorJson(-500, "The domain name has expired", JsonData{})
		return
	}

	//获取当前用户
	userInfo, done := c.getUserInfoSigningKey(signingKey)
	if done {
		c.ErrorJson(-500, "userInfo error", JsonData{})
		return
	}

	blance, err := strconv.ParseFloat(userInfo.Data.Balance, 32)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	lock.Lock()

	markets, err := models.GetNamesMarketInNameStatus(name, 1)
	if err != nil {
		lock.Unlock()
		c.ErrorJson(-500, "The domain name has been processed", JsonData{})
		return
	}

	if blance < float64(markets.Offer) {
		lock.Unlock()
		c.ErrorJson(-500, "Keep account funds greater than "+strconv.Itoa(markets.Offer)+" ae", JsonData{})
		return
	}

	fmt.Println("1111111111111111")

	transfer, b := Transfer(strconv.Itoa(markets.Offer), "ak_25BWMx4An9mmQJNPSwJisiENek3bAGadze31Eetj4K4JJC8VQN", signingKey)
	if b {
		c.ErrorJson(-500, "transfer error", JsonData{})
		lock.Unlock()
		return
	}

	th, done := c.getTh(transfer.Data.Tx.Hash)
	if done {
		lock.Unlock()
		return
	}

	height, i2 := c.getBlockHeight()
	if i2 {
		lock.Unlock()
		return
	}

	err = models.UpdateNamesMarketInsOut(name, markets.InOwner, 1, markets.InToken, transfer.Data.Tx.Hash, userInfo.Data.Address, int(height))
	if err != nil {
		lock.Unlock()
		c.ErrorJson(-500, err.Error()+"", JsonData{})
		return
	}
	lock.Unlock()

	c.SuccessJson(th)
}

func (c *MarketInController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	offer, err := c.GetInt("offer", 0)

	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	if offer < 1 || offer > 100000000 {
		c.ErrorJson(-500, "names error offer < 1 || offer > 100000000", JsonData{})
		return
	}
	nameInfo, done := c.getNameInfo(name)
	//域名有问题或者其他错误
	if done {
		c.ErrorJson(-500, "names error", JsonData{})
		return
	}
	//如果域名还没到账户
	if nameInfo.Data.CurrentHeight <= nameInfo.Data.EndHeight {
		c.ErrorJson(-500, "The domain name is still up for grabs", JsonData{})
		return
	}

	//如果域名已经过期
	if nameInfo.Data.CurrentHeight >= nameInfo.Data.OverHeight {
		c.ErrorJson(-500, "The domain name has expired", JsonData{})
		return
	}

	//如果域名已经过期
	if nameInfo.Data.OverHeight-nameInfo.Data.CurrentHeight < 20*24*30 {
		c.ErrorJson(-500, "Please keep domain name valid for more than one month", JsonData{})
		return
	}

	//获取当前用户
	userInfo, done := c.getUserInfoSigningKey(signingKey)
	if done {
		return
	}

	//如果域名不是账户的归属者
	if nameInfo.Data.Owner != userInfo.Data.Address {
		c.ErrorJson(-500, "The domain name does not belong to you ", JsonData{})
		return
	}

	blance, err := strconv.ParseFloat(userInfo.Data.Balance, 32)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	if blance < 1 {
		c.ErrorJson(-500, "Keep account funds greater than 1 ae", JsonData{})
		return
	}

	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	lock.Lock()

	_, err = models.GetNamesMarketIn(name, userInfo.Data.Address, 0, utils.Md5V(signingKey))
	if err == nil {
		lock.Unlock()
		c.ErrorJson(-500, name+" Waiting for block confirmation", JsonData{})
		return
	}

	transfer, b := c.TransferName(name, signingKey, "ak_25BWMx4An9mmQJNPSwJisiENek3bAGadze31Eetj4K4JJC8VQN")
	if b {
		lock.Unlock()
		return
	}

	th, done := c.getTh(transfer.Data.Hash)
	if done {
		lock.Unlock()
		return
	}

	height, i2 := c.getBlockHeight()
	if i2 {
		lock.Unlock()
		return
	}

	_, err = models.InsertNamesMarketIn(name, offer, userInfo.Data.Address, transfer.Data.Hash, int(height), utils.Md5V(signingKey))
	if err != nil {
		lock.Unlock()
		c.ErrorJson(-500, err.Error()+"", JsonData{})
		return
	}
	lock.Unlock()

	c.SuccessJson(th)
}

func (c *MarketCancelController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")

	nameInfo, done := c.getNameInfo(name)
	//域名有问题或者其他错误
	if done {
		c.ErrorJson(-500, "names error", JsonData{})
		return
	}
	//如果域名还没到账户
	if nameInfo.Data.CurrentHeight <= nameInfo.Data.EndHeight {
		c.ErrorJson(-500, "The domain name is still up for grabs", JsonData{})
		return
	}

	//如果域名已经过期
	if nameInfo.Data.CurrentHeight >= nameInfo.Data.OverHeight {
		c.ErrorJson(-500, "The domain name has expired", JsonData{})
		return
	}

	//获取当前用户
	userInfo, done := c.getUserInfoSigningKey(signingKey)
	if done {
		return
	}

	blance, err := strconv.ParseFloat(userInfo.Data.Balance, 32)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	if blance < 1 {
		c.ErrorJson(-500, "Keep account funds greater than 1 ae", JsonData{})
		return
	}

	lock.Lock()

	_, err = models.GetNamesMarketIn(name, userInfo.Data.Address, 1, utils.Md5V(signingKey))
	if err != nil {
		lock.Unlock()
		c.ErrorJson(-500, "The domain name does not belong to you ", JsonData{})
		return
	}

	aensSigningKey := beego.AppConfig.String("names::signingKey")
	transfer, b := c.TransferName(name, aensSigningKey, userInfo.Data.Address)
	if b {
		lock.Unlock()
		return
	}

	th, done := c.getTh(transfer.Data.Hash)
	if done {
		lock.Unlock()
		return
	}

	height, i2 := c.getBlockHeight()
	if i2 {
		lock.Unlock()
		return
	}

	err = models.UpdateNamesMarketInsCanecl(name, userInfo.Data.Address, 1, utils.Md5V(signingKey), transfer.Data.Hash, int(height))
	if err != nil {
		lock.Unlock()
		c.ErrorJson(-500, err.Error()+"", JsonData{})
		return
	}
	lock.Unlock()
	c.SuccessJson(th)
}
