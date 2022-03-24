package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Menu struct {
		models.BaseModel

		ParentId       uint64 `gorm:"type:bigint unsigned;not null;index;comment:上级ID"`
		Name           string `gorm:"type:varchar(120);not null;comment:菜单名称"`
		Icon           string `gorm:"type:varchar(120);null;comment:图标"`
		Uri            string `gorm:"type:varchar(255);not null;comment:跳转链接"`
		IsLink         bool   `gorm:"type:boolean;default 0;comment:是否链接"`
		PermissionName string `gorm:"type:varchar(60);null;comment:权限名称"`
		GuardName      string `gorm:"type:varchar(30);not null;comment:项目名称"`
		Sequence       uint64 `gorm:"type:int unsigned;null;comment:排序"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Menu{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Menu{})
	}

	migrate.Add("2022_03_24_105051_create_menus_table", up, down)
}
