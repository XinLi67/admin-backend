package seeders

import (
	"fmt"
	"gohub/app/models/permission"
	"gohub/app/models/permission_group"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedPermissionsTable", func(db *gorm.DB) {

		// 管理员权限组
		user := permission_group.GetByName("管理员")
		role := permission_group.GetByName("角色")
		mypermission := permission_group.GetByName("权限")
		permissionGroup := permission_group.GetByName("权限组")
		menu := permission_group.GetByName("菜单")
		common := permission_group.GetByName("公共")

		permissions := []permission.Permission{
			{
				PermissionGroupId: user.ID,
				Name:              "admin-user.index",
				GuardName:         "admin",
				DisplayName:       "管理员列表",
				Icon:              "",
			},
			{
				PermissionGroupId: user.ID,
				Name:              "admin-user.show",
				GuardName:         "admin",
				DisplayName:       "详细",
				Icon:              "",
			},
			{
				PermissionGroupId: user.ID,
				Name:              "admin-user.store",
				GuardName:         "admin",
				DisplayName:       "添加",
				Icon:              "",
			},
			{
				PermissionGroupId: user.ID,
				Name:              "admin-user.update",
				GuardName:         "admin",
				DisplayName:       "修改",
				Icon:              "",
			},
			{
				PermissionGroupId: role.ID,
				Name:              "admin-user.roles",
				GuardName:         "admin",
				DisplayName:       "用户角色列表",
				Icon:              "",
			},
			{
				PermissionGroupId: role.ID,
				Name:              "admin-user.assign-roles",
				GuardName:         "admin",
				DisplayName:       "分配角色",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.index",
				GuardName:         "admin",
				DisplayName:       "角色列表",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.show",
				GuardName:         "admin",
				DisplayName:       "角色详情",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.show",
				GuardName:         "admin",
				DisplayName:       "角色详情",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.store",
				GuardName:         "admin",
				DisplayName:       "新建角色",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.update",
				GuardName:         "admin",
				DisplayName:       "修改角色",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.destroy",
				GuardName:         "admin",
				DisplayName:       "删除角色",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.permissions",
				GuardName:         "admin",
				DisplayName:       "角色权限列表",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.assign-permissions",
				GuardName:         "admin",
				DisplayName:       "角色分配权限",
				Icon:              "",
			}, {
				PermissionGroupId: role.ID,
				Name:              "role.guard-name-roles",
				GuardName:         "admin",
				DisplayName:       "项目角色列表",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.index",
				GuardName:         "admin",
				DisplayName:       "权限列表",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.index",
				GuardName:         "admin",
				DisplayName:       "权限列表",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.show",
				GuardName:         "admin",
				DisplayName:       "权限详情",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.store",
				GuardName:         "admin",
				DisplayName:       "新建权限",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.update",
				GuardName:         "admin",
				DisplayName:       "更新权限",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.destroy",
				GuardName:         "admin",
				DisplayName:       "删除权限",
				Icon:              "",
			}, {
				PermissionGroupId: mypermission.ID,
				Name:              "permission.destroy",
				GuardName:         "admin",
				DisplayName:       "删除权限",
				Icon:              "",
			}, {
				PermissionGroupId: menu.ID,
				Name:              "menu.index",
				GuardName:         "admin",
				DisplayName:       "菜单列表",
				Icon:              "",
			}, {
				PermissionGroupId: menu.ID,
				Name:              "menu.show",
				GuardName:         "admin",
				DisplayName:       "菜单详情",
				Icon:              "",
			}, {
				PermissionGroupId: menu.ID,
				Name:              "menu.store",
				GuardName:         "admin",
				DisplayName:       "新建菜单",
				Icon:              "",
			}, {
				PermissionGroupId: menu.ID,
				Name:              "menu.update",
				GuardName:         "admin",
				DisplayName:       "更新菜单",
				Icon:              "",
			}, {
				PermissionGroupId: menu.ID,
				Name:              "menu.destroy",
				GuardName:         "admin",
				DisplayName:       "删除菜单",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.index",
				GuardName:         "admin",
				DisplayName:       "权限组列表",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.show",
				GuardName:         "admin",
				DisplayName:       "权限组详情",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.store",
				GuardName:         "admin",
				DisplayName:       "新增权限组",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.update",
				GuardName:         "admin",
				DisplayName:       "编辑权限组",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.destroy",
				GuardName:         "admin",
				DisplayName:       "删除权限组",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.guard-name-for-permission",
				GuardName:         "admin",
				DisplayName:       "获取看守器权限",
				Icon:              "",
			}, {
				PermissionGroupId: permissionGroup.ID,
				Name:              "permission-group.all",
				GuardName:         "admin",
				DisplayName:       "所有权限组",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "channel.index",
				GuardName:         "admin",
				DisplayName:       "渠道列表",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "channel.show",
				GuardName:         "admin",
				DisplayName:       "渠道详情",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "channel.store",
				GuardName:         "admin",
				DisplayName:       "新建渠道",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "channel.update",
				GuardName:         "admin",
				DisplayName:       "更新渠道",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "channel.destroy",
				GuardName:         "admin",
				DisplayName:       "删除渠道",
				Icon:              "",
			},
			{
				PermissionGroupId: common.ID,
				Name:              "department.index",
				GuardName:         "admin",
				DisplayName:       "部门列表",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "department.show",
				GuardName:         "admin",
				DisplayName:       "部门详情",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "department.store",
				GuardName:         "admin",
				DisplayName:       "新建部门",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "department.update",
				GuardName:         "admin",
				DisplayName:       "更新部门",
				Icon:              "",
			}, {
				PermissionGroupId: common.ID,
				Name:              "department.destroy",
				GuardName:         "admin",
				DisplayName:       "删除部门",
				Icon:              "",
			},
		}

		result := db.Table("permissions").Create(&permissions)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
