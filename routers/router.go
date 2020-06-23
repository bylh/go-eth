package routers

import (
	"github.com/gin-contrib/cors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"go-eth/middleware/jwt"
	"go-eth/pkg/export"
	"go-eth/pkg/qrcode"
	"go-eth/pkg/upload"
	"go-eth/routers/api"
	"go-eth/routers/api/news"
	"go-eth/routers/api/trade"
	"go-eth/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://*.perceive.top"},
	//	//AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	//AllowAllOrigins: true,
	//	//AllowOriginFunc: func(origin string) bool {
	//	//	return origin == "*"
	//	//},
	//	MaxAge: 12 * time.Hour,
	//}))
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AllowCredentials = true
	//r.Use(cors.New(config))
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	r.GET("/trade/test", trade.Test)
	r.GET("/newsTags", news.GetNewsTags)
	r.GET("/news", news.GetNews)
	// https://github.com/gin-gonic/gin/issues/686
	//apiProxy := r.Group("/hub")
	//apiProxy.Use()
	//{
	//
	//}
	// 注意：此处为完全代理，GET方法的相对路径会直接追加到代理地址，所以内部要处理，进行替换，或者忽略
	//r.GET("/get-hub-type", proxy.ReverseProxy("https://tophub.fun:8080", "/GetType"))
	//r.GET("/netdata", proxy.ReverseProxy("http://bylh.top:19999", ""))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}
	return r
}
