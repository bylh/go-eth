package routers

import (
	"net/http"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

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
	// 已经在nginx配置过服务了，此处不再配置，否则重复设置为报错
	//r.Use(cors.New(cors.Config{：：
	//	AllowOrigins: []string{"https://*.perceive.top", "https://*.bylh.top"},
	//	//AllowMethods:     []string{"PUT", "PATCH"},
	//	//AllowHeaders:  []string{"Origin"},
	//	//ExposeHeaders: []string{"Content-Length"},
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
	mode := gin.Mode()
	if mode == gin.DebugMode || mode == gin.TestMode {
		// 生产环境的cors是在nginx配置的，测试环境要打开
		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"http://local.bylh.top:3000", "http://local.perceive.top:3000", "https://perceive.top", "http://localhost*"},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production
			Debug: true,
		}))
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	r.GET("/trade/getOpenOrders", trade.GetOpenOrders)
	r.GET("/trade/getAllTickers", trade.GetAllTickers)

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

	//store := cookie.NewStore([]byte("secret"))
	// TODO 同一域名及其子域名可以，不同域名后端不可以为前端设置cookie
	//store.Options(sessions.Options{
	//	Domain: ".perceive.top",
	//	Path: "/",
	//	Secure: true,
	//	SameSite: 2,
	//
	//})
	//sessionNames := []string{"a", "b"}
	//apiTrade := r.Group("/session")
	//apiTrade.Use(sessions.SessionsMany(sessionNames, store))
	//{
	//	apiTrade.GET("/login", func(c *gin.Context) {
	//		sessionA := sessions.DefaultMany(c, "a")
	//		sessionB := sessions.DefaultMany(c, "b")
	//		fmt.Println("sessionA", sessionA.Get("a"))
	//		fmt.Println("sessionB", sessionA.Get("b"))
	//
	//		sessionA.Set("hello", "world!")
	//		err := sessionA.Save()
	//		fmt.Println("sessionAERR", err)
	//
	//		if sessionA.Get("hello") != "world!" {
	//			sessionA.Set("hello", "world!")
	//			sessionA.Save()
	//		}
	//		sessionB.Set("hello", "world?")
	//		err = sessionB.Save()
	//		fmt.Println("sessionBERR", err)
	//		if sessionB.Get("hello") != "world?" {
	//			sessionB.Set("hello", "world?")
	//			sessionB.Save()
	//		}
	//
	//		c.JSON(200, gin.H{
	//			"a": sessionA.Get("hello"),
	//			"b": sessionB.Get("hello"),
	//		})
	//	})
	//}

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
