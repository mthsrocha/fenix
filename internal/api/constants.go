package api


var assets_fan_token_gaming = []string{ "NAVI", "OG", "ALL", "TH" }

var assets_fan_token_motorsport = []string{ "SAUBER", "AM" }

var assets_fan_token_outros = []string{ "PFL" }

var assets_crypto = []string{
	"BTC",
	"ETH",
	"LTC",
	"BCH",
	"USDC",
	"XRP",
	"PAXG",
}

var assets_fan_token = []string{
	"ACMFT",
	"SAUBER",
	"ASR",
	"ARG",
	"AM",
	"ATM",
	"BAR",
	"SCCP",
	"PSG",
	"CITY",
	"GAL",
	"NAVI",
	"CAI",
	"JUV",
	"OG",
	"ALL",
	"TH",
	"PFL",
}

var assets_utility_token = []string{
	"ANKR",
	"AXS",
	"BAND",
	"BAT",
	"LINK",
	"CHZ",
	"MANA",
	"ENJ",
	"MCO2",
	"OMG",
	"MATIC",
	"GRT",
	"WBTC",
	"WBX",
}

var assets_digital_assets = []string{
	"MBCONS01",
	"MBCONS02",
	"MBPRK01",
	"MBPRK02",
	"MBPRK03",
	"MBPRK04",
	"MBPRK05",
	"MBVASCO01",
}

var assets_home = []string{
	"BTC",
	"BTC",
	"LTC",
	"BCH",
	"CHZ",
	"XRP",
	"PAXG",
	"MCO2",
	"PSG",
	"CITY",
	"SCCP",
	"NAVI",
}

var assets_fan_token_futebol = []string{
	"ACMFT",
	"SAUBER",
	"ASR",
	"ARG",
	"ATM",
	"BAR",
	"SCCP",
	"PSG",
	"CITY",
	"GAL",
	"CAI",
	"JUV",
}

var all_assets = []string{
	"AAVE",
	"ACMFT",
	"ACORDO01",
	"ALLFT",
	"AMFT",
	"ANKR",
	"ARGFT",
	"ASRFT",
	"ATMFT",
	"AXS",
	"BAL",
	"BAND",
	"BARFT",
	"BAT",
	"BCH",
	"BNT",
	"BTC",
	"CAIFT",
	"CHZ",
	"CITYFT",
	"COMP",
	"CRV",
	"DAI",
	"ENJ",
	"ETH",
	"GALFT",
	"GALOFT",
	"GRT",
	"IMOB01",
	"IMOB02",
	"JUVFT",
	"KNC",
	"LINK",
	"LTC",
	"MANA",
	"MATIC",
	"MBCONS01",
	"MBCONS02",
	"MBFP01",
	"MBFP02",
	"MBFP03",
	"MBFP04",
	"MBFP05",
	"MBPRK01",
	"MBPRK02",
	"MBPRK03",
	"MBPRK04",
	"MBPRK05",
	"MBVASCO01",
	"MCO2",
	"MKR",
	"NAVIFT",
	"OGFT",
	"OMG",
	"PAXG",
	"PFLFT",
	"PSGFT",
	"REN",
	"SAUBERFT",
	"SCCPFT",
	"SNX",
	"SUSHI",
	"THFT",
	"UMA",
	"UNI",
	"USDC",
	"WBTC",
	"WBX",
	"XRP",
	"YBOFT",
	"YFI",
	"ZRX",
}

func GetHomeAssets() []string {
	return assets_home
}

func GetCryptos() []string {
	return assets_crypto
}

func GetFanTokens() []string {
	return assets_fan_token
}

func GetUtilityTokens() []string {
	return assets_utility_token
}

func GetDigitalAssets() []string {
	return assets_digital_assets
}

func GetFanTokensFutebol() []string {
	return assets_fan_token_futebol
}

func GetFanTokensGaming() []string {
	return assets_fan_token_gaming
}

func GetFanTokensMotorsports() []string {
	return assets_fan_token_motorsport
}