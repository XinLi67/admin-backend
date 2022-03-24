package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type UserRole struct {
		models.BaseModel

		UserId uint64 `gorm:"type:bigint unsigned;not null;index;comment:用户ID"`
		RoleId uint64 `gorm:"type:bigint unsigned;not null;index;comment:角色ID"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&UserRole{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&UserRole{})
	}

	migrate.Add("2022_03_24_112451_create_user_roles_table", up, down)
}
