package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-eth/go-copy/models"
	"github.com/go-eth/go-copy/pkg/gredis"
	"github.com/go-eth/go-copy/pkg/logging"
	"github.com/go-eth/go-copy/pkg/setting"
	"github.com/go-eth/go-copy/routers"
	"github.com/go-eth/go-copy/pkg/util"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/go-eth/go-copy
// @license.name MIT
// @license.url https://github.com/go-eth/go-copy/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

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
