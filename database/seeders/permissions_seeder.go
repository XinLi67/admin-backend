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

		userGroup := permission_group.GetByName("管理员")
		permissions := []permission.Permission{
			{
				PermissionGroupId: userGroup.ID,
				Name:              "admin-user.index",
				GuardName:         "admin",
				DisplayName:       "管理员列表",
				Icon:              "",
			},
			{
				PermissionGroupId: userGroup.ID,
				Name:              "admin-user.show",
				GuardName:         "admin",
				DisplayName:       "详细",
				Icon:              "",
			},
			{
				PermissionGroupId: userGroup.ID,
				Name:              "admin-user.store",
				GuardName:         "admin",
				DisplayName:       "添加",
				Icon:              "",
			},
			{
				PermissionGroupId: userGroup.ID,
				Name:              "admin-user.update",
				GuardName:         "admin",
				DisplayName:       "修改",
				Icon:              "",
			},
			{
				PermissionGroupId: userGroup.ID,
				Name:              "admin-user.destroy",
				GuardName:         "admin",
				DisplayName:       "详细",
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
