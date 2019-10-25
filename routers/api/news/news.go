package news

import (
	"github.com/gin-gonic/gin"
	"go-eth/pkg/app"
	"go-eth/pkg/e"
	"go-eth/service/news_service"
	"net/http"
)

func GetNews(c *gin.Context) {
	appG := app.Gin{C: c}
	// data, _ := news_service.fetchNews();
	news_service.FetchNews()
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}