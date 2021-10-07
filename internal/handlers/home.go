package handlers

import (
	"net/http"

	"github.com/Fenix/internal/controllers"
)

func HomeHandlers() {
	http.HandleFunc("/", controllers.HomeTemplate)
	http.HandleFunc("/home/", controllers.HomeTemplate)

	http.HandleFunc("/home/cryptocoins/prices/", controllers.CriptocoinsPricesTemplate)
	http.HandleFunc("/home/cryptocoins/btc/", controllers.CriptocoinsBTCTemplate)
	http.HandleFunc("/home/cryptocoins/bch/", controllers.CriptocoinsBCHTemplate)
	http.HandleFunc("/home/cryptocoins/eth/", controllers.CriptocoinsETHTemplate)
	http.HandleFunc("/home/cryptocoins/ltc/", controllers.CriptocoinsLTCTemplate)
	http.HandleFunc("/home/cryptocoins/paxg/", controllers.CriptocoinsPAXGTemplate)
	http.HandleFunc("/home/cryptocoins/xrp/", controllers.CriptocoinsXRPTemplate)
	http.HandleFunc("/home/cryptocoins/about/", controllers.CriptocoinsAboutTemplate)
	http.HandleFunc("/home/cryptocoins/others/", controllers.CriptocoinsOthersTemplate)

	http.HandleFunc("/home/digitalassets/prices/", controllers.DigitalAssetsPricesTemplate)
	http.HandleFunc("/home/digitalassets/mbcons01/", controllers.DigitalAssetsMBCONS01Template)
	http.HandleFunc("/home/digitalassets/mbprk03/", controllers.DigitalAssetsMBPRK01Template)
	http.HandleFunc("/home/digitalassets/mbprk04/", controllers.DigitalAssetsMBPRK03Template)
	http.HandleFunc("/home/digitalassets/mbprk01/", controllers.DigitalAssetsMBPRK04Template)
	http.HandleFunc("/home/digitalassets/about/", controllers.DigitalAssetsAboutTemplate)
	http.HandleFunc("/home/digitalassets/others/", controllers.DigitalAssetsOthersTemplate)

	http.HandleFunc("/home/fantokens/prices/", controllers.FanTokensPricesTemplate)
	http.HandleFunc("/home/fantokens/futebol/", controllers.FanTokensFutebolTemplate)
	http.HandleFunc("/home/fantokens/gaming/", controllers.FanTokensGamingTemplate)
	http.HandleFunc("/home/fantokens/motorsports/", controllers.FanTokensMotorsportsTemplate)
	http.HandleFunc("/home/fantokens/about/", controllers.FanTokensAboutTemplate)
	http.HandleFunc("/home/fantokens/others/", controllers.FanTokensOthersTemplate)

	http.HandleFunc("/home/utilitytokens/prices/", controllers.UtilityTokensPricesTemplate)
	http.HandleFunc("/home/utilitytokens/chz/", controllers.UtilityTokensCHZTemplate)
	http.HandleFunc("/home/utilitytokens/mco2/", controllers.UtilityTokensMCO2Template)
	http.HandleFunc("/home/utilitytokens/wbx/", controllers.UtilityTokensWBXTemplate)
	http.HandleFunc("/home/utilitytokens/about/", controllers.UtilityTokensAboutTemplate)
	http.HandleFunc("/home/utilitytokens/others/", controllers.UtilityTokensOthersTemplate)
}
