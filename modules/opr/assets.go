package opr

// V1Assets is the list of assets PegNet launched with
var V1Assets = []string{
	"PNT",
	"USD",
	"EUR",
	"JPY",
	"GBP",
	"CAD",
	"CHF",
	"INR",
	"SGD",
	"CNY",
	"HKD",
	"KRW",
	"BRL",
	"PHP",
	"MXN",
	"XAU",
	"XAG",
	"XPD",
	"XPT",
	"XBT",
	"ETH",
	"LTC",
	"RVN",
	"XBC",
	"FCT",
	"BNB",
	"XLM",
	"ADA",
	"XMR",
	"DASH",
	"ZEC",
	"DCR",
}

// V2Assets contains the following changes to V1:
//		* Rename PNT to PEG
//		* Drop XPD
// 		* Drop XPT
var V2Assets = []string{
	"PEG",
	"USD",
	"EUR",
	"JPY",
	"GBP",
	"CAD",
	"CHF",
	"INR",
	"SGD",
	"CNY",
	"HKD",
	"KRW",
	"BRL",
	"PHP",
	"MXN",
	"XAU",
	"XAG",
	"XBT",
	"ETH",
	"LTC",
	"RVN",
	"XBC",
	"FCT",
	"BNB",
	"XLM",
	"ADA",
	"XMR",
	"DASH",
	"ZEC",
	"DCR",
}

type AssetFloat struct {
	Name  string
	Value float64
}
type AssetUint struct {
	Name  string
	Value uint64
}