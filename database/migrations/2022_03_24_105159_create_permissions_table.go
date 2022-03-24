package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Permission struct {
		models.BaseModel

		PermissionGroupId uint64 `gorm:"type:bigint unsigned;not null;index;comment:权限组ID"`
		Name              string `gorm:"type:varchar(120);not null;comment:权限名称"`
		Icon              string `gorm:"type:varchar(120);null;comment:图标"`
		GuardName         string `gorm:"type:varchar(30);not null;comment:项目名称"`
		DisplayName       string `gorm:"type:varchar(30);not null;comment:显示名称"`
		Description       string `gorm:"type:varchar(255);null;comment:权限说明"`
		Sequence          uint64 `gorm:"type:int unsigned;null;comment:排序"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Permission{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Permission{})
	}

	migrate.Add("2022_03_24_105159_create_permissions_table", up, down)
}
