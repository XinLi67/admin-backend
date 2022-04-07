package seeders

import (
	"fmt"
	"gohub/app/models/menu"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedMenusTable", func(db *gorm.DB) {

		menus := []menu.Menu{
			{
				ParentId:       0,
				Icon:           "Orange",
				Uri:            "/dashboard",
				IsLink:         false,
				PermissionName: "",
				Name:           "仪表盘",
				GuardName:      "admin",
				Sequence:       0,
			}, {
				ParentId:       0,
				Icon:           "Orange",
				Uri:            "/channel",
				IsLink:         false,
				PermissionName: "",
				Name:           "渠道管理",
				GuardName:      "admin",
				Sequence:       0,
			}, {
				ParentId:       0,
				Icon:           "Setting",
				Uri:            "/admin",
				IsLink:         false,
				PermissionName: "",
				Name:           "系统管理",
				GuardName:      "admin",
				Sequence:       0,
			}, {
				ParentId:       0,
				Icon:           "",
				Uri:            "/advertising",
				IsLink:         false,
				PermissionName: "",
				Name:           "广告管理",
				GuardName:      "admin",
				Sequence:       0,
			},
		}

		result := db.Table("menus").Create(&menus)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))

		systemMenu := menu.GetByName("系统管理")

		systemMenus := []menu.Menu{
			{
				ParentId:       systemMenu.ID,
				Icon:           "",
				Uri:            "/admin-user",
				IsLink:         false,
				PermissionName: "",
				Name:           "用户管理",
				GuardName:      "admin",
				Sequence:       0,
			}, {
				ParentId:       systemMenu.ID,
				Icon:           "",
				Uri:            "/department",
				IsLink:         false,
				PermissionName: "",
				Name:           "部门管理",
				GuardName:      "admin",
				Sequence:       0,
			},
			{
				ParentId:       systemMenu.ID,
				Icon:           "",
				Uri:            "/role",
				IsLink:         false,
				PermissionName: "",
				Name:           "角色管理",
				GuardName:      "admin",
				Sequence:       0,
			},
			{
				ParentId:       systemMenu.ID,
				Icon:           "",
				Uri:            "/permission",
				IsLink:         false,
				PermissionName: "",
				Name:           "权限管理",
				GuardName:      "admin",
				Sequence:       0,
			},
			{
				ParentId:       systemMenu.ID,
				Icon:           "",
				Uri:            "/menu",
				IsLink:         false,
				PermissionName: "",
				Name:           "菜单管理",
				GuardName:      "admin",
				Sequence:       0,
			},
		}

		systemResult := db.Table("menus").Create(&systemMenus)

		if err := systemResult.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", systemResult.Statement.Table, systemResult.RowsAffected))

		advertisingMenu := menu.GetByName("广告管理")

		advertisingMenus := []menu.Menu{
			{
				ParentId:       advertisingMenu.ID,
				Icon:           "",
				Uri:            "/advertising-index",
				IsLink:         false,
				PermissionName: "",
				Name:           "广告列表",
				GuardName:      "admin",
				Sequence:       0,
			},
		}

		advertisingResult := db.Table("menus").Create(&advertisingMenus)

		if err := advertisingResult.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", advertisingResult.Statement.Table, advertisingResult.RowsAffected))
	})
}
