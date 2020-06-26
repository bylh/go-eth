package trade_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/binance"
	"github.com/nntaoli-project/goex/builder"
	"github.com/nntaoli-project/goex/okex"
	"go-eth/pkg/setting"
	"os"
	"time"
)

type ExtendGoex interface {
	goex.API
	GetAllUnfinishOrders() ([]goex.Order, error)
}

var apiBuilder = builder.NewAPIBuilder().HttpTimeout(5 * time.Second)

func GetAllUnfinishOrders(i ExtendGoex) ([]goex.Order, error) {
	return i.GetAllUnfinishOrders()
}
func Test(conditions map[string]string) (map[string]interface{}, error) {
	//environ := os.Environ()
	//for i := range environ {
	//	fmt.Println("hah:", i, environ[i])
	//}
	fmt.Println("**************************")
	env := os.Getenv("ENV")
	fmt.Printf("GOPATH is %s\n", env)
	//build spot api
	//api := apiBuilder.APIKey("").APISecretkey("").ClientID("123").Build(goex.BITSTAMP)
	//api := apiBuilder.APIKey(key).APISecretkey(secret).Build(goex.BINANCE)
	// log.Println(api.GetExchangeName())
	// log.Println(api.GetTicker(goex.BTC_USD))
	// log.Println(api.GetDepth(2, goex.BTC_USD))
	//log.Println(api.GetAccount())
	//log.Println(api.GetUnfinishOrders(goex.BTC_USD))

	//build future api
	// futureApi := apiBuilder.APIKey("").APISecretkey("").BuildFuture(goex.HBDM)
	// log.Println(futureApi.GetExchangeName())
	// log.Println(futureApi.GetFutureTicker(goex.BTC_USD, goex.QUARTER_CONTRACT))
	// log.Println(futureApi.GetFutureDepth(goex.BTC_USD, goex.QUARTER_CONTRACT, 5))
	//log.Println(futureApi.GetFutureUserinfo()) // account
	//log.Println(futureApi.GetFuturePosition(goex.BTC_USD , goex.QUARTER_CONTRACT))//position info

	/* ------------------------- ticker ------------------------------- */
	//ticker, err := api.GetTicker(goex.BTC_USD)
	//fmt.Println("BTC_USD", ticker)
	// 可扩展的交易对 NAS
	//ticker, err := api.GetTicker(goex.CurrencyPair{CurrencyA: goex.Currency{Symbol: "NAS", Desc: "https://nebulas.io/cn/"}, CurrencyB: goex.BTC})
	//fmt.Println("BTC_NAS", ticker, err)

	//orders, err := api.GetUnfinishOrders(goex.CurrencyPair{CurrencyA: goex.Currency{Symbol: "NAS", Desc: "https://nebulas.io/cn/"}, CurrencyB: goex.BTC})
	//fmt.Println("GetUnfinishOrders", orders, len(orders))

	result, err := GetOpenOrders(conditions)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]interface{})
	resultMap[conditions["exName"]] = result
	return resultMap, nil

}

/* ------------------------- 根据交易所名称获取交易所api ------------------------------- */
func getExApi(exName string) goex.API {
	var key, secret string
	var exBase string

	switch exName {
	case "BINANCE":
		key = setting.TradeSetting.BinanceKey
		secret = setting.TradeSetting.BinanceSecret
		exBase = goex.BINANCE
		break
	case "OKEX":
		key = setting.TradeSetting.OkexKey
		secret = setting.TradeSetting.OkexSecret
		exBase = goex.OKEX
	}

	println("gin mode", gin.Mode())
	mode := gin.Mode()
	if mode == gin.DebugMode || mode == gin.TestMode {
		apiBuilder = apiBuilder.HttpProxy("socks5://127.0.0.1:7891")
	}
	return apiBuilder.APIKey(key).APISecretkey(secret).Build(exBase)
}

/* ------------------------- 根据交易所名称获取交易所进行中的订单（挂单） ------------------------------- */
func GetOpenOrders(conditions map[string]string) ([]goex.Order, error) {
	//var api interface{}
	api := getExApi(conditions["exName"])
	base := conditions["base"]
	target := conditions["target"]
	if len(base) == 0 || len(target) == 0 {
		switch conditions["exName"] {
		case "BINANCE":
			//var b = binance.Binance{setting.TradeSetting.BinanceKey,
			//	setting.TradeSetting.BinanceSecret,
			//
			//}
			// 可以引入不同交易所的实例，使用同样的http配置
			var b = binance.New(apiBuilder.GetHttpClient(), setting.TradeSetting.BinanceKey, setting.TradeSetting.BinanceSecret)
			//accessKey: setting.TradeSetting.BinanceKey,
			//secretKey: setting.TradeSetting.BinanceSecret,
			//baseUrl:   goex.BINANCE,
			return b.GetAllUnfinishOrders()
		}
		//test := ExtendGoex(api)
		//return GetAllUnfinishOrders(test)
	}
	return api.GetUnfinishOrders(goex.CurrencyPair{CurrencyA: goex.Currency{Symbol: target, Desc: ""}, CurrencyB: goex.Currency{Symbol: base, Desc: ""}})
}
func GetAllTickers() (*[]goex.FutureTicker, error) {
	api := getExApi("OKEX")
	if api == nil {
		return nil, nil
	}
	okexF := okex.NewOKExSwap(&goex.APIConfig{
		HttpClient:   apiBuilder.GetHttpClient(),
		ApiKey:       setting.TradeSetting.OkexKey,
		ApiSecretKey: setting.TradeSetting.OkexSecret,
		Endpoint:     "https://www.okex.com",
	})
	return okexF.GetFutureAllTicker()
	//accessKey: setting.TradeSetting.BinanceKey,
	//secretKey: setting.TradeSetting.BinanceSecret,
	//baseUrl:   goex.BINANCE,
}
func GetTicker(target string, base string) (*goex.Ticker, error) {
	// 暂时使用币安的
	api := getExApi("BINANCE")
	return api.GetTicker(goex.CurrencyPair{CurrencyA: goex.Currency{Symbol: target}, CurrencyB: goex.Currency{Symbol: base}})
}
func GetAccount(exName string) (*goex.Account, error) {
	api := getExApi(exName)
	return api.GetAccount()
}
