package main

/*
包前是下划线_：当导入一个包时，该包下的文件里所有init函数都会被执行，但是有时我们仅仅需要使用init函数而已并不希望把整个包导入（不使用包里的其他函数）

包前是点.：

import（.“fmt”）
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
*/

import (
	"fmt"
	_ "go-eth/docs"
	"go-eth/service/news_service"
	"log"

	// 如果使用诸如 http.StatusOK 之类的常量，则需要引入 net/http 包
	"net/http"

	// 设置https

	"go-eth/models"
	"go-eth/pkg/gredis"
	"go-eth/pkg/logging"
	"go-eth/pkg/setting"
	"go-eth/pkg/util"
	"go-eth/routers"
)

/*
 bylh:
 init()和main()是go语言中的保留函数，两个函数在go语言中的区别如下
 1、两个函数在定义时不能有任何的参数和返回值
 2、该函数只能由go程序自动调用不可以被引用
 3、init可以应用于任意包中，且可以重复定义多个，main函数只能应用于main包中，且只能定义一个
 4、对同一个go文件的init(), 调用顺序是从上往下的
 5、对同一个 package 中的不同文件，将文件名按字符串进行“从小到大”排序，之后顺序调用各文件中的init()函数
 6、对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数
 7、如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，一次执行对应的 init 方法。
 8、在同一个文件中，常量、变量、init()、main() 依次进行初始化。

*/
/* 初始化 */
func init() {
	fmt.Println("开始初始化项目")
	// 读取配置文件， 初始化基本配置
	setting.Setup()
	// 根据配置文件建立数据库连接
	models.Setup()
	// 根据配置设置log文件
	logging.Setup()
	// 根据配置文件设置redis TODO 错误处理
	gredis.Setup()
	// 根据配置文件中app->JwtSecret设置鉴权
	util.Setup()

	// 消息接口轮询任务
	news_service.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/go-eth
// @license.name MIT
// @license.url https://go-eth/blob/master/LICENSE

func main() {
	// 根据配置文件中server字段设置 端口 读写超时等配置
	//gin.SetMode(setting.ServerSetting.RunMode)
	//mode := gin.Mode()
	//if mode == gin.DebugMode || mode == gin.TestMode {
	//	可以在这里读取不同的配置（启动或构建时指定ginMode:  GIN_MODE=release go run main.go）
	//}

	// 初始化项目路由信息
	routersInit := routers.InitRouter()
	// 设置读写超时时间
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout

	// Addr    string  // TCP address to listen on, ":http" if empty
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort) // endPoint为 :8001, 库中要求的string

	// << n 左移运算符即乘以2的n次方， >> n 右移运算符即除以2的n次方 次数为1 * 2^20
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,    // 监听地址
		Handler:        routersInit, // 路由等
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	// serve开始监听 TODO 错误处理
	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
