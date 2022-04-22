// Package routes 注册路由
package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"gohub/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	r.StaticFS("/public", http.Dir("./public"))
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
	v1.Use(middlewares.LimitIP("10000-H"))

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
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("1000-H"), vcc.ShowCaptcha)

		}

		// 权限组
		pg := new(controllers.PermissionGroupsController)
		v1.GET("permission-group-all", pg.All)
		pgGroup := v1.Group("/permission-group")
		{
			pgGroup.GET("", pg.Index)
			pgGroup.POST("", middlewares.AuthJWT(), pg.Store)
			pgGroup.POST("/:id/update", middlewares.AuthJWT(), pg.Update)
			pgGroup.POST("/:id/delete", middlewares.AuthJWT(), pg.Delete)
			pgGroup.POST("/batch-delete", middlewares.AuthJWT(), pg.BatchDelete)
			pgGroup.GET("/:id", pg.Show)
		}

		// 权限
		permission := new(controllers.PermissionsController)
		v1.GET("permission-user-all", permission.All)
		permissionGroup := v1.Group("/permission")
		{
			permissionGroup.GET("", permission.Index)
			permissionGroup.POST("", middlewares.AuthJWT(), permission.Store)
			permissionGroup.POST("/:id/update", middlewares.AuthJWT(), permission.Update)
			permissionGroup.POST("/:id/delete", middlewares.AuthJWT(), permission.Delete)
			permissionGroup.POST("/batch-delete", middlewares.AuthJWT(), permission.BatchDelete)
			permissionGroup.GET("/:id", permission.Show)
		}

		// 菜单管理
		mu := new(controllers.MenusController)
		v1.GET("my-menu", mu.MyMenu)
		menu := v1.Group("/menu")
		{
			menu.GET("", mu.Index)
			menu.POST("", middlewares.AuthJWT(), mu.Store)
			menu.POST("/:id/update", middlewares.AuthJWT(), mu.Update)
			menu.POST("/:id/delete", middlewares.AuthJWT(), mu.Delete)
			menu.POST("/batch-delete", middlewares.AuthJWT(), mu.BatchDelete)
			menu.GET("/:id", mu.Show)
		}

		// 角色管理
		role := new(controllers.RolesController)
		v1.GET("guard-name-roles/:guardName", role.GetGuardNameRoles)
		roleGroup := v1.Group("/role")
		{
			roleGroup.GET("", role.Index)
			roleGroup.POST("", middlewares.AuthJWT(), role.Store)
			roleGroup.POST("/:id/update", middlewares.AuthJWT(), role.Update)
			roleGroup.POST("/:id/delete", middlewares.AuthJWT(), role.Delete)
			roleGroup.POST("/batch-delete", middlewares.AuthJWT(), role.BatchDelete)
			roleGroup.GET("/:id", role.Show)
			roleGroup.POST("/:id/roles/:guardName", role.UpdateGuardName)
			roleGroup.GET("/:id/permissions", role.GetRolePermissions)
			roleGroup.POST("/:id/permissions", role.AssignPermissions)
		}

		uc := new(controllers.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/admin-user")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.POST("", middlewares.AuthJWT(), uc.Store)
			usersGroup.POST("/:id/update", middlewares.AuthJWT(), uc.Update)
			usersGroup.POST("/:id/delete", middlewares.AuthJWT(), uc.Delete)
			usersGroup.POST("/batch-delete", middlewares.AuthJWT(), uc.BatchDelete)
			usersGroup.GET("/:id", uc.Show)
			usersGroup.POST("/password-reset", middlewares.AuthJWT(), uc.UpdatePassword)
		}

		// 部门管理
		dep := new(controllers.DepartmentsController)
		departmentGroup := v1.Group("/department")
		{
			departmentGroup.GET("", dep.Index)
			departmentGroup.POST("", middlewares.AuthJWT(), dep.Store)
			departmentGroup.POST("/:id/update", middlewares.AuthJWT(), dep.Update)
			departmentGroup.POST("/:id/delete", middlewares.AuthJWT(), dep.Delete)
			departmentGroup.POST("/batch-delete", middlewares.AuthJWT(), dep.BatchDelete)
			departmentGroup.GET("/:id", dep.Show)
		}

		// 渠道管理
		channel := new(controllers.ChannelsController)
		channelGroup := v1.Group("/channel")
		{
			channelGroup.GET("", channel.Index)
			channelGroup.POST("", middlewares.AuthJWT(), channel.Store)
			channelGroup.POST("/:id/update", middlewares.AuthJWT(), channel.Update)
			channelGroup.POST("/:id/delete", middlewares.AuthJWT(), channel.Delete)
			channelGroup.POST("/batch-delete", middlewares.AuthJWT(), channel.BatchDelete)
			channelGroup.GET("/:id", channel.Show)
		}

		// 广告管理
		advertising := new(controllers.AdvertisingsController)
		advertisingGroup := v1.Group("/advertising")
		{
			advertisingGroup.GET("", advertising.Index)
			advertisingGroup.POST("", middlewares.AuthJWT(), advertising.Store)
			advertisingGroup.POST("/:id/update", middlewares.AuthJWT(), advertising.Update)
			advertisingGroup.POST("/:id/delete", middlewares.AuthJWT(), advertising.Delete)
			advertisingGroup.POST("/batch-delete", middlewares.AuthJWT(), advertising.BatchDelete)
			advertisingGroup.GET("/:id", advertising.Show)
			advertisingGroup.GET("/advertising-position-id/:id", advertising.IndexByAdvertisingPosId)
			advertisingGroup.GET("/export", advertising.Export)
		}

		// 公告管理
		announcement := new(controllers.AnnouncementsController)
		announcementGroup := v1.Group("/announcement")
		{
			announcementGroup.GET("", announcement.Index)
			announcementGroup.POST("", middlewares.AuthJWT(), announcement.Store)
			announcementGroup.POST("/:id/update", middlewares.AuthJWT(), announcement.Update)
			announcementGroup.POST("/:id/delete", middlewares.AuthJWT(), announcement.Delete)
			announcementGroup.POST("/batch-delete", middlewares.AuthJWT(), announcement.BatchDelete)
			announcementGroup.GET("/:id", announcement.Show)
			announcementGroup.GET("/export", announcement.Export)
		}

		// 广告位管理
		advertisingPosition := new(controllers.AdvertisingPositionsController)
		advertisingPositionGroup := v1.Group("/advertising-position")
		{
			advertisingPositionGroup.GET("", advertisingPosition.Index)
			advertisingPositionGroup.POST("", middlewares.AuthJWT(), advertisingPosition.Store)
			advertisingPositionGroup.POST("/:id/update", middlewares.AuthJWT(), advertisingPosition.Update)
			advertisingPositionGroup.POST("/:id/delete", middlewares.AuthJWT(), advertisingPosition.Delete)
			advertisingPositionGroup.POST("/batch-delete", middlewares.AuthJWT(), advertisingPosition.BatchDelete)
			advertisingPositionGroup.GET("/:id", advertisingPosition.Show)
			advertisingPositionGroup.GET("/export", advertisingPosition.Export)
		}

		// 公告位管理
		announcementPosition := new(controllers.AnnouncementPositionsController)
		announcementPositionGroup := v1.Group("/announcement-position")
		{
			announcementPositionGroup.GET("", announcementPosition.Index)
			announcementPositionGroup.POST("", middlewares.AuthJWT(), announcementPosition.Store)
			announcementPositionGroup.POST("/:id/update", middlewares.AuthJWT(), announcementPosition.Update)
			announcementPositionGroup.POST("/:id/delete", middlewares.AuthJWT(), announcementPosition.Delete)
			announcementPositionGroup.POST("/batch-delete", middlewares.AuthJWT(), announcementPosition.BatchDelete)
			announcementPositionGroup.GET("/:id", announcementPosition.Show)
			announcementPositionGroup.GET("/export", announcementPosition.Export)
		}

		// 广告计划管理
		advertisingPlan := new(controllers.AdvertisingPlansController)
		advertisingPlanGroup := v1.Group("/advertising-plan")
		{
			advertisingPlanGroup.GET("", advertisingPlan.Index)
			advertisingPlanGroup.POST("", middlewares.AuthJWT(), advertisingPlan.Store)
			advertisingPlanGroup.POST("/:id/update", middlewares.AuthJWT(), advertisingPlan.Update)
			advertisingPlanGroup.POST("/:id/delete", middlewares.AuthJWT(), advertisingPlan.Delete)
			advertisingPlanGroup.POST("/batch-delete", middlewares.AuthJWT(), advertisingPlan.BatchDelete)
			advertisingPlanGroup.POST("/batch-store", middlewares.AuthJWT(), advertisingPlan.BatchStore)
			advertisingPlanGroup.GET("/:id", advertisingPlan.Show)
			advertisingPlanGroup.GET("/export", advertisingPlan.Export)
		}

		// 公告计划管理
		announcementPlan := new(controllers.AnnouncementPlansController)
		announcementPlanGroup := v1.Group("/announcement-plan")
		{
			announcementPlanGroup.GET("", announcementPlan.Index)
			announcementPlanGroup.POST("", middlewares.AuthJWT(), announcementPlan.Store)
			announcementPlanGroup.POST("/:id/update", middlewares.AuthJWT(), announcementPlan.Update)
			announcementPlanGroup.POST("/:id/delete", middlewares.AuthJWT(), announcementPlan.Delete)
			announcementPlanGroup.POST("/batch-delete", middlewares.AuthJWT(), announcementPlan.BatchDelete)
			announcementPlanGroup.POST("/batch-store", middlewares.AuthJWT(), announcementPlan.BatchStore)
			announcementPlanGroup.GET("/:id", announcementPlan.Show)
			announcementPlanGroup.GET("/export", announcementPlan.Export)
		}

		// 素材管理
		material := new(controllers.MaterialsController)
		materialGroup := v1.Group("/material")
		{
			materialGroup.GET("", material.Index)
			materialGroup.POST("", middlewares.AuthJWT(), material.Store)
			materialGroup.POST("/:id/update", middlewares.AuthJWT(), material.Update)
			materialGroup.POST("/:id/delete", middlewares.AuthJWT(), material.Delete)
			materialGroup.POST("/batch-delete", middlewares.AuthJWT(), material.BatchDelete)
			materialGroup.GET("/:id", material.Show)

		}

		// 素材组管理
		mg := new(controllers.MaterialGroupsController)
		mgGroup := v1.Group("/material-group")
		{
			mgGroup.GET("", mg.Index)
			mgGroup.POST("", middlewares.AuthJWT(), mg.Store)
			mgGroup.POST("/:id/update", middlewares.AuthJWT(), mg.Update)
			mgGroup.POST("/:id/delete", middlewares.AuthJWT(), mg.Delete)
			mgGroup.POST("/batch-delete", middlewares.AuthJWT(), mg.BatchDelete)
			mgGroup.GET("/:id", mg.Show)
			mgGroup.GET("/:id/document", mg.GetDocumentById)
			mgGroup.GET("/:id/menu", mg.GetTree)
		}

		// 点击记录管理
		clickRecord := new(controllers.ClickRecordsController)
		clickRecordGroup := v1.Group("/click-record")
		{
			clickRecordGroup.GET("", clickRecord.Index)
			clickRecordGroup.POST("", middlewares.AuthJWT(), clickRecord.Store)
			clickRecordGroup.POST("/:id/update", middlewares.AuthJWT(), clickRecord.Update)
			clickRecordGroup.POST("/:id/delete", middlewares.AuthJWT(), clickRecord.Delete)
			clickRecordGroup.POST("/batch-delete", middlewares.AuthJWT(), clickRecord.BatchDelete)
			clickRecordGroup.GET("/:id", clickRecord.Show)
			clickRecordGroup.GET("/showByMonth", clickRecord.ShowByMonth)
			clickRecordGroup.GET("/showByWeek", clickRecord.ShowByWeek)
			clickRecordGroup.GET("/showByAdversingId", clickRecord.ShowByAdversingId)
			clickRecordGroup.GET("/showByCustomerId", clickRecord.ShowByCustomerId)
			clickRecordGroup.GET("/showByPosId", clickRecord.ShowByPosId)
		}

		// 审核记录管理
		auditRecord := new(controllers.AuditRecordsController)
		auditRecordGroup := v1.Group("/audit-record")
		{
			auditRecordGroup.GET("", auditRecord.Index)
			auditRecordGroup.POST("", middlewares.AuthJWT(), auditRecord.Store)
			auditRecordGroup.POST("/:id/update", middlewares.AuthJWT(), auditRecord.Update)
			auditRecordGroup.POST("/:id/delete", middlewares.AuthJWT(), auditRecord.Delete)
			auditRecordGroup.POST("/batch-delete", middlewares.AuthJWT(), auditRecord.BatchDelete)
			auditRecordGroup.GET("/:id", auditRecord.Show)
		}

		tester := new(controllers.TestersController)
		v1.POST("/tester", tester.BatchDelete)

		//上传文件接口
		uploadController := new((controllers.UploadController))
		upload := v1.Group("/upload")
		{
			upload.POST("", uploadController.Upload)
		}

	}
}
