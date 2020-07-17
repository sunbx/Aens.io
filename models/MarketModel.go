package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type NamesMarket struct {
	ID           int    `orm:"column(id)" json:"id"`
	Name         string `orm:"column(name)" json:"name"`
	Offer        int    `orm:"column(offer)" json:"offer"`
	InOwner      string `orm:"column(in_owner)" json:"in_owner"`
	InStatus     int    `orm:"column(in_status)" json:"in_status"`
	InHeight     int    `orm:"column(in_height)" json:"in_height"`
	InTx         string `orm:"column(in_tx)" json:"in_tx"`
	InTime       int64  `orm:"column(in_time)" json:"in_time"`
	InToken      string `orm:"column(in_token)" json:"in_token"`
	CancelTx     string `orm:"column(cancel_tx)" json:"cancel_tx"`
	CancelStatus int    `orm:"column(cancel_status)" json:"cancel_status"`
	CancelHeight int    `orm:"column(cancel_height)" json:"cancel_height"`
	CancelTime   int64  `orm:"column(cancel_time)" json:"cancel_time"`
	OutOwner     string `orm:"column(out_owner)" json:"out_owner"`
	OutTx        string `orm:"column(out_tx)" json:"out_tx"`
	OutStatus    int    `orm:"column(out_status)" json:"out_status"`
	OutHeight    int    `orm:"column(out_height)" json:"out_height"`
	OutTime      int64  `orm:"column(out_time)" json:"out_time"`
	TxName       string `orm:"column(tx_name)" json:"tx_name"`
	TxToken      string `orm:"column(tx_token)" json:"tx_token"`
	Time         int64  `orm:"column(time)" json:"time"`
	Height       int    `orm:"column(height)" json:"height"`
}

func (n *NamesMarket) TableName() string {
	return "names_market"
}

func RegisterNamesMarket() {
	orm.RegisterModel(new(NamesMarket))
}

func InsertNamesMarketIn(name string, offer int, inOwner string, inTx string, inHeight int, inToken string) (int64, error) {
	if name == "" {
		return -1, errors.Errorf("name null")
	}
	if inOwner == "" {
		return -1, errors.Errorf("inOwner null")
	}
	if inTx == "" {
		return -1, errors.Errorf("inTx null")
	}
	if inToken == "" {
		return -1, errors.Errorf("inToken null")
	}
	if inHeight <= 0 {
		return -1, errors.Errorf("inHeight error")
	}
	if offer <= 0 || offer > 10000000 {
		return -1, errors.Errorf("offer offer <= 0 || offer > 10000000")
	}

	unix := time.Now().UnixNano() / 1e6
	namesMarket := NamesMarket{
		Name:     name,
		Offer:    offer,
		InOwner:  inOwner,
		InTx:     inTx,
		InHeight: inHeight,
		InTime:   unix,
		InToken:  inToken,
	}
	id, err := orm.NewOrm().Insert(&namesMarket)
	return id, err
}

func GetNamesMarketIn(name string, inOwner string, inStatus int, inToken string) (NamesMarket, error) {
	var namesMarket NamesMarket
	err := orm.NewOrm().QueryTable("names_market").
		Filter("name", name).
		Filter("in_owner", inOwner).
		Filter("in_status", inStatus).
		Filter("in_token", inToken).
		Filter("cancel_tx", "").
		Filter("out_tx", "").
		Filter("tx_name", "").
		Filter("tx_token", "").
		One(&namesMarket)
	return namesMarket, err
}

func GetNamesMarketInNameStatus(name string, inStatus int) (NamesMarket, error) {
	var namesMarket NamesMarket
	err := orm.NewOrm().QueryTable("names_market").
		Filter("name", name).
		Filter("in_status", inStatus).
		Filter("cancel_tx", "").
		Filter("out_tx", "").
		Filter("tx_name", "").
		Filter("tx_token", "").
		One(&namesMarket)
	return namesMarket, err
}

func GetNamesMarketIns(inOwner string, inStatus int) ([]NamesMarket, error) {
	var namesMarket []NamesMarket
	_, err := orm.NewOrm().QueryTable("names_market").
		Filter("in_owner", inOwner).
		Filter("in_status", inStatus).
		Filter("cancel_tx", "").
		Filter("out_tx", "").
		Filter("tx_name", "").
		Filter("tx_token", "").
		All(&namesMarket)
	return namesMarket, err
}

func GetNamesMarketInsStaus(inStatus int) ([]NamesMarket, error) {
	var namesMarket []NamesMarket
	_, err := orm.NewOrm().QueryTable("names_market").
		Filter("in_status", inStatus).
		Filter("cancel_tx", "").
		Filter("out_tx", "").
		Filter("tx_name", "").
		Filter("tx_token", "").
		All(&namesMarket)
	return namesMarket, err
}

func GetNamesMarketOutStaus() ([]NamesMarket, error) {
	var namesMarket []NamesMarket
	_, err := orm.NewOrm().QueryTable("names_market").
		Filter("in_status", 1).
		Filter("cancel_tx", "").
		Filter("tx_name", "").
		Filter("tx_token", "").
		All(&namesMarket)
	return namesMarket, err
}

func GetNamesMarketMarketStaus() ([]NamesMarket, error) {
	var namesMarket []NamesMarket
	_, err := orm.NewOrm().QueryTable("names_market").
		Filter("in_status", 1).
		Filter("out_status", 1).
		Filter("cancel_tx", "").
		All(&namesMarket)
	return namesMarket, err
}

func UpdateNamesMarketInsStaus(id int, inStatus int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("id", id).Update(orm.Params{
		"in_status": inStatus,
	})
	return err
}

func UpdateNamesMarketOutsStaus(id int, outStatus int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("id", id).Update(orm.Params{
		"out_status": outStatus,
	})
	return err
}

func UpdateNamesMarketOutsDefIn(id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("id", id).Update(orm.Params{
		"out_tx":     "",
		"out_height": 0,
		"out_owner":  "",
		"out_time":   0,
	})
	return err
}

func UpdateNamesMarketNameTx(id int, txName string, height int) error {
	unix := time.Now().UnixNano() / 1e6
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("id", id).Update(orm.Params{
		"tx_name": txName,
		"time":    unix,
		"height":  height,
	})
	return err
}

func UpdateNamesMarketTokenTx(id int, txToken string, height int) error {
	unix := time.Now().UnixNano() / 1e6
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("id", id).Update(orm.Params{
		"tx_token": txToken,
		"time":     unix,
		"height":   height,
	})
	return err
}

func UpdateNamesMarketInsCanecl(name string, inOwner string, inStatus int, inToken string, cancelTx string, cancelHeight int) error {
	unix := time.Now().UnixNano() / 1e6
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("name", name).Filter("in_owner", inOwner).Filter("in_status", inStatus).Filter("in_token", inToken).Update(orm.Params{
		"in_status":     2,
		"cancel_tx":     cancelTx,
		"cancel_status": 1,
		"cancel_height": cancelHeight,
		"cancel_time":   unix,
	})
	return err
}

func UpdateNamesMarketInsOut(name string, inOwner string, inStatus int, inToken string, outTx string, outOwner string, outHeight int) error {
	unix := time.Now().UnixNano() / 1e6
	o := orm.NewOrm()
	_, err := o.QueryTable("names_market").Filter("name", name).Filter("in_owner", inOwner).Filter("in_status", inStatus).Filter("in_token", inToken).Update(orm.Params{
		"out_tx":     outTx,
		"out_height": outHeight,
		"out_owner":  outOwner,
		"out_time":   unix,
	})
	return err
}
