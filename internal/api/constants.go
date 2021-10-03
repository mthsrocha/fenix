package api

var assets_fan_token_gaming = []string{"NAVI", "OG", "ALL", "TH"}

var assets_fan_token_motorsport = []string{"SAUBER", "AM"}

var assets_fan_token_outros = []string{"PFL"}

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
	"SAUBERFT",
	"ASRFT",
	"ARGFT",
	"AMFT",
	"ATMFT",
	"BARFT",
	"SCCPFT",
	"PSGFT",
	"CITYFT",
	"GALFT",
	"NAVIFT",
	"CAIFT",
	"JUVFT",
	"OGFT",
	"ALLFT",
	"THFT",
	"PFLFT",
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
	"ETH",
	"LTC",
	"BCH",
	"CHZ",
	"XRP",
	"PAXG",
	"MCO2",
	"PSGFT",
	"CITYFT",
	"SCCPFT",
	"NAVIFT",
}

var assets_fan_token_futebol = []string{
	"ACMFT",
	"ASRFT",
	"ARGFT",
	"ATMFT",
	"BARFT",
	"SCCPFT",
	"PSGFT",
	"CITYFT",
	"GALFT",
	"CAIFT",
	"JUVFT",
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
