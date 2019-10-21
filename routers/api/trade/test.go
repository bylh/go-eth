package trade

import (
	"github.com/gin-gonic/gin"
	"go-eth/pkg/app"
	"go-eth/pkg/e"
	"go-eth/pkg/setting"
	"go-eth/service/trade_service"
	"net/http"
)

func Test(c *gin.Context) {
	appG := app.Gin{C: c}
	data, _ := trade_service.Test(setting.TradeSetting.HuobiKey, setting.TradeSetting.HuobiSecret)
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
