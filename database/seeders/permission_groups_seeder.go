package seeders

import (
	"fmt"
	"gohub/app/models/permission_group"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedPermissionGroupsTable", func(db *gorm.DB) {

		permissionGroups := []permission_group.PermissionGroup{
			{
				Name:        "管理员",
				Description: "",
			},
			{
				Name:        "角色",
				Description: "",
			},
			{
				Name:        "权限",
				Description: "",
			},
			{
				Name:        "权限组",
				Description: "",
			},
			{
				Name:        "菜单",
				Description: "",
			},
			{
				Name:        "公共",
				Description: "",
			},
		}

		result := db.Table("permission_groups").Create(&permissionGroups)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
