package routers

import (
	"github.com/astaxie/beego"
	"names/controllers"
)

func init() {
	//静态页

	beego.Router("/", &controllers.HomeController{})
	//切换语言
	beego.Router("/language", &controllers.LanguageController{})

	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/login/logout", &controllers.LoginLogoutController{})

	beego.Router("/auction", &controllers.AuctionController{})

	beego.Router("/detail/address", &controllers.DetailAddressController{})

	beego.Router("/my/register", &controllers.AuctionMyController{})

	beego.Router("/my/over", &controllers.ExpireMyController{})

	beego.Router("/price", &controllers.PriceController{})

	beego.Router("/expire", &controllers.ExpireController{})

	beego.Router("/detail", &controllers.DetailController{})

	beego.Router("/create", &controllers.CreateController{})

	beego.Router("/otc/hosting", &controllers.OTCHostingController{})

	beego.Router("/otc/market", &controllers.OTCMarketController{})

	beego.Router("/api/name/info", &controllers.ApiNamesInfoController{})

	beego.Router("/api/user/info", &controllers.ApiUserInfoController{})

	beego.Router("/api/otc/market/in", &controllers.MarketInController{})

	beego.Router("/api/otc/market/cancel", &controllers.MarketCancelController{})

	beego.Router("/api/otc/market/out", &controllers.MarketOutController{})

}
