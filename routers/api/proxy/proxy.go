package proxy

import (
	"github.com/gin-gonic/gin"
	"go-eth/pkg/logging"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(target string, replacePath string) gin.HandlerFunc {
	uri, err := url.Parse(target)
	if err != nil {
		logging.Warn(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(uri)
	return func(c *gin.Context) {
		c.Request.URL.Path = replacePath
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

//func ReverseProxy() gin.HandlerFunc {
//
//	target := "tophub.fun:8080/GetType"
//
//	return func(c *gin.Context) {
//		director := func(req *http.Request) {
//			//r := c.Request
//			//req = r
//			req.URL.Scheme = "http"
//			req.URL.Host = target
//		}
//		proxy := &httputil.ReverseProxy{Director: director}
//		proxy.ServeHTTP(c.Writer, c.Request)
//	}
//}

//func ReverseProxy() gin.HandlerFunc {
//
//	target := "localhost:3000"
//
//	return func(c *gin.Context) {
//		director := func(req *http.Request) {
//			r := c.Request
//			req = r
//			req.URL.Scheme = "http"
//			req.URL.Host = target
//			req.Header["my-header"] = []string{r.Header.Get("my-header")}
//			// Golang camelcases headers
//			delete(req.Header, "My-Header")
//		}
//		proxy := &httputil.ReverseProxy{Director: director}
//		proxy.ServeHTTP(c.Writer, c.Request)
//	}
//}
