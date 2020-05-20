package controllers

import (
	"fmt"
	"names/utils"
	"strconv"
	"time"
)

type HomeController struct {
	BaseController
}
type RegisterController struct {
	BaseController
}

type LoginController struct {
	BaseController
}

type LoginLogoutController struct {
	BaseController
}
type PriceController struct {
	BaseController
}
type AuctionController struct {
	BaseController
}
type AuctionMyController struct {
	BaseController
}
type ExpireMyController struct {
	BaseController
}

type ExpireController struct {
	BaseController
}

type DetailController struct {
	BaseController
}

type CreateController struct {
	BaseController
}

func (c *HomeController) Get() {

	nameBase, done := c.getNameBase()
	if done {
		return
	}
	c.Data["sum_price"] = nameBase.Data.SumPrice
	c.Data["sum"] = nameBase.Data.Sum
	c.Data["count"] = nameBase.Data.Count
	for i := 0; i < len(nameBase.Data.Ranking); i++ {
		nameBase.Data.Ranking[i].Id = i + 1
		nameBase.Data.Ranking[i].Address = nameBase.Data.Ranking[i].Owner[0:3] + "****" + nameBase.Data.Ranking[i].Owner[len(nameBase.Data.Ranking[i].Owner)-5:]
	}
	c.Data["ranking"] = nameBase.Data.Ranking
	address := c.Ctx.GetCookie("address")
	if address == "" {
		c.Data["address"] = "Not logged in, click login"
		c.Data["login_text"] = "Login"
	} else {
		c.Data["address"] = address[0:3] + "****" + address[len(address)-5:]
		c.Data["login_text"] = "Logout"
	}

	fmt.Println("address->", address)
	fmt.Println("namebase->", nameBase)
	c.TplName = "index.html"
}

func (c *LoginLogoutController) Get() {
	address := c.Ctx.GetCookie("address")
	if address == "" {
		c.Redirect("/login", 302)
	} else {
		c.Ctx.SetCookie("address", "")
		c.Redirect("/", 302)
	}
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}
func (c *AuctionController) Get() {
	page, _ := c.GetInt("page", 1)
	name, done := c.getAuctionName(strconv.Itoa(page))
	if done {
		return
	}
	for i := 0; i < len(name.Data); i++ {
		name.Data[i].Id = (i + 1) + (page-1)*20
		name.Data[i].Address = name.Data[i].Owner[0:3] + "****" + name.Data[i].Owner[len(name.Data[i].Owner)-5:]
		name.Data[i].Progress = int64(float64(name.Data[i].CurrentHeight-name.Data[i].StartHeight) / (float64(name.Data[i].EndHeight - name.Data[i].StartHeight)) * 100)
		name.Data[i].Countdown = utils.StrTime(time.Now().Unix() - (name.Data[i].EndHeight-name.Data[i].CurrentHeight)*3*60)
		cPrice, _ := strconv.ParseFloat(name.Data[i].CurrentPrice, 64)
		price, _ := strconv.ParseFloat(name.Data[i].Price, 64)
		name.Data[i].Gains = price / cPrice * 100
	}

	c.Data["name"] = name.Data
	c.Data["pageLeftDisplay"] = "display: block"
	c.Data["pageRightDisplay"] = "display: block"
	if page-1 < 1 {
		c.Data["pageLeftDisplay"] = "display: none"
	}
	if len(name.Data) < 20 {
		c.Data["pageRightDisplay"] = "display: none"
	}
	c.Data["pageLeft"] = page - 1
	c.Data["pageRight"] = page + 1

	c.TplName = "auction.html"
}

func (c *AuctionMyController) Get() {
	page, _ := c.GetInt("page", 1)
	address := c.Ctx.GetCookie("address")
	if address != "" {
		name, done := c.getAuctionMyName(strconv.Itoa(page), address)
		if done {
			return
		}
		for i := 0; i < len(name.Data); i++ {
			name.Data[i].Id = (i + 1) + (page-1)*20
			name.Data[i].Address = name.Data[i].Owner[0:3] + "****" + name.Data[i].Owner[len(name.Data[i].Owner)-5:]
			name.Data[i].Progress = int64(float64(name.Data[i].CurrentHeight-name.Data[i].StartHeight) / (float64(name.Data[i].EndHeight - name.Data[i].StartHeight)) * 100)
			name.Data[i].Countdown = utils.StrTime(time.Now().Unix() - (name.Data[i].EndHeight-name.Data[i].CurrentHeight)*3*60)
			cPrice, _ := strconv.ParseFloat(name.Data[i].CurrentPrice, 64)
			price, _ := strconv.ParseFloat(name.Data[i].Price, 64)
			name.Data[i].Gains = price / cPrice * 100
		}

		c.Data["name"] = name.Data
		c.Data["pageLeftDisplay"] = "display: block"
		c.Data["pageRightDisplay"] = "display: block"
		if page-1 < 1 {
			c.Data["pageLeftDisplay"] = "display: none"
		}
		if len(name.Data) < 20 {
			c.Data["pageRightDisplay"] = "display: none"
		}
		c.Data["pageLeft"] = page - 1
		c.Data["pageRight"] = page + 1
		c.TplName = "my_auction.html"
	} else {
		c.Redirect("/login", 302)
	}

}

func (c *PriceController) Get() {
	page, _ := c.GetInt("page", 1)
	name, done := c.getPriceName(strconv.Itoa(page))
	if done {
		return
	}
	for i := 0; i < len(name.Data); i++ {
		name.Data[i].Id = (i + 1) + (page-1)*20
		name.Data[i].Address = name.Data[i].Owner[0:3] + "****" + name.Data[i].Owner[len(name.Data[i].Owner)-5:]
		name.Data[i].Progress = int64(float64(name.Data[i].CurrentHeight-name.Data[i].StartHeight) / (float64(name.Data[i].EndHeight - name.Data[i].StartHeight)) * 100)
		name.Data[i].Countdown = utils.StrTime(time.Now().Unix() - (name.Data[i].EndHeight-name.Data[i].CurrentHeight)*3*60)

		cPrice, _ := strconv.ParseFloat(name.Data[i].CurrentPrice, 64)
		price, _ := strconv.ParseFloat(name.Data[i].Price, 64)
		name.Data[i].Gains = utils.Decimal((cPrice - price) / price * 100)
	}

	c.Data["name"] = name.Data
	c.Data["pageLeftDisplay"] = "display: block"
	c.Data["pageRightDisplay"] = "display: block"
	if page-1 < 1 {
		c.Data["pageLeftDisplay"] = "display: none"
	}
	if len(name.Data) < 20 {
		c.Data["pageRightDisplay"] = "display: none"
	}
	c.Data["pageLeft"] = page - 1
	c.Data["pageRight"] = page + 1

	c.TplName = "price.html"
}
func (c *ExpireController) Get() {
	page, _ := c.GetInt("page", 1)

	name, done := c.getExpireName(strconv.Itoa(page))
	if done {
		return
	}
	for i := 0; i < len(name.Data); i++ {
		name.Data[i].Id = (i + 1) + (page-1)*20
		name.Data[i].Address = name.Data[i].Owner[0:3] + "****" + name.Data[i].Owner[len(name.Data[i].Owner)-5:]
		name.Data[i].Progress = int64(float64(name.Data[i].CurrentHeight-(name.Data[i].OverHeight-50000)) / (float64(name.Data[i].OverHeight - (name.Data[i].OverHeight - 50000))) * 100)
		name.Data[i].Countdown = utils.StrTime(time.Now().Unix() - (name.Data[i].OverHeight-name.Data[i].CurrentHeight)*3*60)

		cPrice, _ := strconv.ParseFloat(name.Data[i].CurrentPrice, 64)
		price, _ := strconv.ParseFloat(name.Data[i].Price, 64)
		name.Data[i].Gains = utils.Decimal((cPrice - price) / price * 100)
	}

	c.Data["name"] = name.Data
	c.Data["pageLeftDisplay"] = "display: block"
	c.Data["pageRightDisplay"] = "display: block"
	if page-1 < 1 {
		c.Data["pageLeftDisplay"] = "display: none"
	}
	if len(name.Data) < 20 {
		c.Data["pageRightDisplay"] = "display: none"
	}
	c.Data["pageLeft"] = page - 1
	c.Data["pageRight"] = page + 1

	c.TplName = "expire.html"
}
func (c *ExpireMyController) Get() {
	page, _ := c.GetInt("page", 1)
	address := c.Ctx.GetCookie("address")
	if address != "" {
		name, done := c.getRegisterMyName(strconv.Itoa(page), address)
		if done {
			return
		}
		for i := 0; i < len(name.Data); i++ {
			name.Data[i].Id = (i + 1) + (page-1)*20
			name.Data[i].Address = name.Data[i].Owner[0:3] + "****" + name.Data[i].Owner[len(name.Data[i].Owner)-5:]
			name.Data[i].Progress = int64(float64(name.Data[i].CurrentHeight-(name.Data[i].OverHeight-50000)) / (float64(name.Data[i].OverHeight - (name.Data[i].OverHeight - 50000))) * 100)
			name.Data[i].Countdown = utils.StrTime(time.Now().Unix() - (name.Data[i].OverHeight-name.Data[i].CurrentHeight)*3*60)

			cPrice, _ := strconv.ParseFloat(name.Data[i].CurrentPrice, 64)
			price, _ := strconv.ParseFloat(name.Data[i].Price, 64)
			name.Data[i].Gains = utils.Decimal((cPrice - price) / price * 100)
		}

		c.Data["name"] = name.Data
		c.Data["pageLeftDisplay"] = "display: block"
		c.Data["pageRightDisplay"] = "display: block"
		if page-1 < 1 {
			c.Data["pageLeftDisplay"] = "display: none"
		}
		if len(name.Data) < 20 {
			c.Data["pageRightDisplay"] = "display: none"
		}
		c.Data["pageLeft"] = page - 1
		c.Data["pageRight"] = page + 1

		c.TplName = "my_expire.html"
	} else {
		c.Redirect("/login", 302)
	}

}

func (c *DetailController) Get() {
	address := c.Ctx.GetCookie("address")
	if address == "" {
		c.Redirect("/login", 302)
		return
	}
	name := c.GetString("name")
	nameInfo, done := c.getNameInfo(name)
	if done {
		return
	}
	for i := 0; i < len(nameInfo.Data.Claim); i++ {
		cPrice, _ := strconv.ParseFloat(nameInfo.Data.Claim[i].NameFee, 64)
		price, _ := strconv.ParseFloat(nameInfo.Data.Price, 64)
		nameInfo.Data.Claim[i].Gains = utils.Decimal(((cPrice / price) - 1) * 100)
		nameInfo.Data.Claim[i].Address = nameInfo.Data.Claim[i].AccountID[0:3] + "****" + nameInfo.Data.Claim[i].AccountID[len(nameInfo.Data.Claim[i].AccountID)-5:]
		t := time.Unix(nameInfo.Data.Claim[i].Time/1000, 0)
		dateStr := t.Format("2006/01/02 15:04:05")
		nameInfo.Data.Claim[i].Date = dateStr
	}
	c.Data["update_display"] = "display: none"
	c.Data["transfer_display"] = "display: none"
	c.Data["claim_display"] = "display: none"

	if nameInfo.Data.CurrentHeight < nameInfo.Data.EndHeight {
		c.Data["claim_display"] = "display: inline-block"
	}

	if nameInfo.Data.Owner == address && nameInfo.Data.CurrentHeight < nameInfo.Data.OverHeight && nameInfo.Data.CurrentHeight > nameInfo.Data.EndHeight {
		c.Data["update_display"] = "display: inline-block"
		c.Data["transfer_display"] = "display: inline-block"
	}

	c.Data["nameInfo"] = nameInfo.Data
	c.TplName = "detail.html"

}

func (c *CreateController) Get() {
	name := c.GetString("name")
	address := c.Ctx.GetCookie("address")
	if address == "" {
		c.Redirect("/login", 302)
		return
	} else {
		info, _ := c.getUserInfo(address)
		c.Data["balance"] = info.Data.Balance
		c.Data["address"] = address

		c.Data["name"] = name
		c.TplName = "create.html"
	}

}
func (c *LoginController) Get() {
	c.TplName = "login.html"
}
