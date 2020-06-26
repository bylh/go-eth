package trade

import (
	"github.com/gin-gonic/gin"
	"go-eth/pkg/app"
	"go-eth/pkg/e"
	"go-eth/service/trade_service"
	"net/http"
)

func GetOpenOrders(c *gin.Context) {
	appG := app.Gin{C: c}
	data, _ := trade_service.GetOpenOrders(map[string]string{"exName": c.Query("exName"), "base": c.Query("base"), "target": c.Query("target")})
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
func GetAllTickers(c *gin.Context) {
	appG := app.Gin{C: c}
	data, _ := trade_service.GetAllTickers()
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
func GetAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	data, _ := trade_service.GetAccount(c.Query("exName"))
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
