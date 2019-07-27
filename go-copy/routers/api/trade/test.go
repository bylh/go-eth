package trade

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eth/go-copy/pkg/app"
	"github.com/go-eth/go-copy/pkg/e"
	"net/http"
	//"github.com/go-eth/go-copy/service/trade_service"
)

func Test(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": "test",
	})
}
