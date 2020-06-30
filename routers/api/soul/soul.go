package soul

import (
	"github.com/gin-gonic/gin"
	"go-eth/pkg/app"
	"go-eth/pkg/e"
	"go-eth/service/soul_service"
	"net/http"
	"strconv"
)

func GetSouls(c *gin.Context) {
	appG := app.Gin{C: c}
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	maps := map[string]interface{}{}
	//maps["tag"] = c.Query("tag")
	data, _ := soul_service.GetSouls(maps, offset, limit)
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
