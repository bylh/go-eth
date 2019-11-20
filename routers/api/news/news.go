package news

import (
	"github.com/gin-gonic/gin"
	"go-eth/pkg/app"
	"go-eth/pkg/e"
	"go-eth/service/news_service"
	"net/http"
	"strconv"
)

func GetNews(c *gin.Context) {
	appG := app.Gin{C: c}
	// data, _ := news_service.fetchNews();
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	maps := map[string]interface{}{}
	maps["tag"] = c.Query("tag")
	data, _ := news_service.GetNews(maps, pageNum, pageSize)
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetNewsTags(c *gin.Context) {
	appG := app.Gin{C: c}
	// data, _ := news_service.fetchNews();
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	maps := map[string]interface{}{}
	data, _ := news_service.GetNewsTags(maps, pageNum, pageSize)
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
