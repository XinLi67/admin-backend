// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	routes.RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},                                                           // 允许的前端地址
			AllowMethods:     []string{"PUT", "GET", "DELETE", "POST", "PATCH", "OPTIONS"},            //允许的方法
			AllowHeaders:     []string{"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token"}, //添加的header
			ExposeHeaders:    []string{"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "*" // 允许的前端地址
			},
			MaxAge: 12 * time.Hour,
		}),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
