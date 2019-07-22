package common

/* 		All the assets on pegnet

PegNet, 				PNT, 		PNT

US Dollar, 				USD, 		pUSD
Euro, 					EUR, 		pEUR
Japanese Yen, 			JPY, 		pJPY
Pound Sterling, 		GBP, 		pGBP
Canadian Dollar, 		CAD,		pCAD
Swiss Franc, 			CHF,		pCHF
Indian Rupee, 			INR,		pINR
Singapore Dollar, 		SGD, 		pSGD
Chinese Yuan, 			CNY,		pCNY
Hong Kong Dollar, 		HKD,		pHKD
Tiawanese Dollar,		TWD, 		pTWD
Korean Won,				KRW,		pKRW
Argentine Peso,			ARS,		pARS
Brazil Real,			BRL,		pBRL
Philippine Peso			PHP,		pPHP
Mexican Peso			MXN,		pMXN

Gold Troy Ounce, 		XAU,		pXAU
Silver Troy Ounce, 		XAG,		pXAG
Palladium Troy Ounce, 	XPD,		pXPD
Platinum Troy Ounce, 	XPT,		pXPT

Bitcoin, 				XBT,		pXBT
Ethereum, 				ETH,		pETH
Litecoin, 				LTC,		pLTC
Ravencoin,				RVN,		pRVN
Bitcoin Cash, 			XBC,		pXBC
Factom, 				FCT,		pFCT
Binance Coin			BNB, 		pBNB
Stellar					XLM, 		pXLM
Cardano					ADA, 		pADA
Monero					XMR, 		pXMR
Dash					DASH,		pDASH
Zcash					ZCASH,		pZCASH
Decred					DCR,		pDCR
*/

var (
	CurrencyAssets = []string{
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
		"TWD",
		"KRW",
		"ARS",
		"BRL",
		"PHP",
		"MXN",
	}

	CommodityAssets = []string{
		"XAU",
		"XAG",
		"XPD",
		"XPT",
	}

	CryptoAssets = []string{
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
		"ZCASH",
		"DCR",
	}

	AllAssets = merge(CurrencyAssets, CommodityAssets, CryptoAssets)
)

func AssetListContains(assetList []string, asset string) bool {
	for _, a := range assetList {
		if asset == a {
			return true
		}
	}
	return false
}

func merge(assets ...[]string) []string {
	acc := []string{}
	for _, list := range assets {
		acc = append(acc, list...)
	}
	return acc
}

var AssetNames = []string{
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
	"XAU",
	"XAG",
	"XPD",
	"XPT",
	"XBT",
	"ETH",
	"LTC",
	"XBC",
	"FCT",
}