package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"names/controllers"
	"names/models"
	_ "names/routers"
	"strconv"
)

//引入数据模型
func init() {
	orm.Debug = false
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")
	dbconnect := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	_ = orm.RegisterDataBase("default", "mysql", dbconnect /*"root:root@tcp(localhost:3306)/test?charset=utf8"*/) //密码为空格式
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	beego.AddFuncMap("i18n", i18n.Tr)

	models.RegisterNamesMarket()

}

func main() {
	task()
	beego.Run()
	task()
}

var isTask = true

func task() {
	tk := toolbox.NewTask("myTask", "0/10 * * * * *", func() error {
		if (isTask) {
			isTask = false
			TaskInsStatus()
			TaskOutStatus()
			TaskMarketStatus()
			isTask = true
		} else {
		}

		return nil
	})

	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
	fmt.Println("定时任务开启")

}

func TaskMarketStatus() {
	markets, err := models.GetNamesMarketMarketStaus()
	if err != nil {
		fmt.Println("TaskInsStatus->", err)
		return
	}
	for i := 0; i < len(markets); i++ {
		if markets[i].OutStatus == 1 {

			height, b2 := controllers.GetBlockHeight()
			if b2 {
				fmt.Println("MarketHeight")
				continue
			}
			aensSigningKey := beego.AppConfig.String("names::signingKey")

			if markets[i].TxToken == ""{
				walletTransfer, i2 := controllers.Transfer(strconv.Itoa(markets[i].Offer), markets[i].InOwner, aensSigningKey)
				if i2 {
					fmt.Println("Market TransferName ERROR")
					continue
				}
				fmt.Println("Market TransferName ->", walletTransfer)
				err = models.UpdateNamesMarketTokenTx(markets[i].ID, walletTransfer.Data.Tx.Hash, int(height))
				if err != nil {
					fmt.Println("UpdateNamesMarketNameTx->", err)
					continue
				}
			}

			if markets[i].TxName == ""{
				transfer, b := controllers.TransferName(markets[i].Name, aensSigningKey, markets[i].OutOwner)
				if b {
					fmt.Println("Market TransferName ERROR")
					continue
				}
				err := models.UpdateNamesMarketNameTx(markets[i].ID, transfer.Data.Hash, int(height))
				if err != nil {
					fmt.Println("UpdateNamesMarketNameTx->", err)
				}
			}
		}
	}
}

func TaskOutStatus() {
	markets, err := models.GetNamesMarketOutStaus()
	if err != nil {
		fmt.Println("TaskInsStatus->", err)
		return
	}
	for i := 0; i < len(markets); i++ {
		if markets[i].OutTx != "" {
			th, b := controllers.GetTh(markets[i].InTx)
			if b {
				fmt.Println("OutTx")
				continue
			}
			height, b2 := controllers.GetBlockHeight()
			if b2 {
				fmt.Println("height")
				continue
			}
			if th.Data.BlockHeight > 0 && height-th.Data.BlockHeight >= 1 {
				_ = models.UpdateNamesMarketOutsStaus(markets[i].ID, 1)
			}
			if th.Data.Hash == "" {
				_ = models.UpdateNamesMarketOutsDefIn(markets[i].ID)
			}
		}
	}
}

func TaskInsStatus() {
	markets, err := models.GetNamesMarketInsStaus(0)
	if err != nil {
		fmt.Println("TaskInsStatus->", err)
		return
	}
	for i := 0; i < len(markets); i++ {
		th, b := controllers.GetTh(markets[i].InTx)
		if b {
			fmt.Println("InTx")
			continue
		}
		height, b2 := controllers.GetBlockHeight()
		if b2 {
			fmt.Println("height")
			continue
		}
		if th.Data.BlockHeight > 0 && height-th.Data.BlockHeight >= 1 {
			_ = models.UpdateNamesMarketInsStaus(markets[i].ID, 1)
		}
		if th.Data.Hash == "" {
			_ = models.UpdateNamesMarketInsStaus(markets[i].ID, 3)
		}
	}
}
