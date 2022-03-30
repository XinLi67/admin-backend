// Package routes 注册路由
package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 注册用户
			// suc := new(auth.SignupController)
			// authGroup.POST("/signup", middlewares.GuestJWT(), suc.Signup)
			// authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			lgc := new(auth.LoginController)
			authGroup.POST("/login", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/logout", middlewares.AuthJWT(), lgc.Logout)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)

		}

		// 权限组
		pg := new(controllers.PermissionGroupsController)
		v1.GET("permission-group-all", pg.All)
		pgGroup := v1.Group("/permission-group")
		{
			pgGroup.GET("", pg.Index)
			pgGroup.POST("", middlewares.AuthJWT(), pg.Store)
			pgGroup.PATCH("/:id", middlewares.AuthJWT(), pg.Update)
			pgGroup.DELETE("/:id", middlewares.AuthJWT(), pg.Delete)
			pgGroup.GET("/:id", pg.Show)
		}

		// 权限
		permission := new(controllers.PermissionsController)
		v1.GET("permission-user-all", permission.All)
		permissionGroup := v1.Group("/permission")
		{
			permissionGroup.GET("", permission.Index)
			permissionGroup.POST("", middlewares.AuthJWT(), permission.Store)
			permissionGroup.PATCH("/:id", middlewares.AuthJWT(), permission.Update)
			permissionGroup.DELETE("/:id", middlewares.AuthJWT(), permission.Delete)
			permissionGroup.GET("/:id", permission.Show)
		}

		// 菜单管理
		mu := new(controllers.MenusController)
		v1.GET("my-menu", mu.MyMenu)
		menu := v1.Group("/menu")
		{
			menu.GET("", mu.Index)
			menu.POST("", middlewares.AuthJWT(), mu.Store)
			menu.PATCH("/:id", middlewares.AuthJWT(), mu.Update)
			menu.DELETE("/:id", middlewares.AuthJWT(), mu.Delete)
			menu.GET("/:id", permission.Show)
		}

		// 角色管理
		role := new(controllers.RolesController)
		v1.GET("guard-name-roles/:guardName", role.GetGuardNameRoles)
		roleGroup := v1.Group("/role")
		{
			roleGroup.GET("", role.Index)
			roleGroup.POST("", middlewares.AuthJWT(), role.Store)
			roleGroup.PATCH("/:id", middlewares.AuthJWT(), role.Update)
			roleGroup.DELETE("/:id", middlewares.AuthJWT(), role.Delete)
			roleGroup.GET("/:id", role.Show)
			roleGroup.PUT("/:id/roles/:guardName", role.UpdateGuardName)
			roleGroup.GET("/:id/permissions", role.GetRolePermissions)
			roleGroup.PUT("/:id/permissions", role.AssignPermissions)
		}

		uc := new(controllers.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/admin-user")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.POST("", middlewares.AuthJWT(), uc.Store)
			usersGroup.PATCH("/:id", middlewares.AuthJWT(), uc.Update)
			usersGroup.DELETE("/:id", middlewares.AuthJWT(), uc.Delete)
			usersGroup.GET("/:id", uc.Show)
		}

		// 部门管理
		dep := new(controllers.DepartmentsController)
		departmentGroup := v1.Group("/department")
		{
			departmentGroup.GET("", dep.Index)
			departmentGroup.POST("", middlewares.AuthJWT(), dep.Store)
			departmentGroup.PATCH("/:id", middlewares.AuthJWT(), dep.Update)
			departmentGroup.DELETE("/:id", middlewares.AuthJWT(), dep.Delete)
			departmentGroup.GET("/:id", dep.Show)
		}

		// 渠道管理
		channel := new(controllers.ChannelsController)
		channelGroup := v1.Group("/channel")
		{
			channelGroup.GET("", channel.Index)
			channelGroup.POST("", middlewares.AuthJWT(), channel.Store)
			channelGroup.PATCH("/:id", middlewares.AuthJWT(), channel.Update)
			channelGroup.DELETE("/:id", middlewares.AuthJWT(), channel.Delete)
			channelGroup.GET("/:id", channel.Show)
		}

		tester := new(controllers.TestersController)
		v1.POST("/tester", tester.BatchDelete)
	}
}
