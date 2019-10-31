package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//NoCache is a middleware function that appends headers
//to prevent the client form caching HTTP request
func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-control", "no-cache,no-store,max-age=0,must-revalidate,value")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		c.Next()
	}
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodOptions {
			c.Next()
		} else {
			//容许跨域
			c.Header("Access-Control-Allow-Origin", "*")
			//客户端所要访问的资源允许使用的方法或方法列表
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			//服务支持的请求方法
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			//在访问中可以使用那些请求头
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			//数据内容类型
			c.Header("Content-Type", "application/json")
			//设置状态码
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		//容许展现页面,安全保护
		c.Header("X-Frame-Options", "DENY")
		//安全设置  script 和 styleSheet 元素会拒绝包含错误的 MIME 类型的响应
		c.Header("X-Content-Type-Options", "nosniff")
		//防止xss攻击设置
		c.Header("X-XSS-Protection", "1; mode=block")
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000")
		}
	}
}
