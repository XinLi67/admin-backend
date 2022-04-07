package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Advertising struct {
		models.BaseModel

		AdvertisingNo         uint64 `gorm:"type:bigint unsigned;column:advertising_no"`
		AdvertisingPositionId uint64 `gorm:"type:bigint unsigned;column:advertising_position_id"`
		CreatorId             uint64 `gorm:"type:bigint unsigned;column:creator_id"`
		DepartmentId          uint64 `gorm:"type:bigint unsigned;column:department_id"`
		Title                 string `gorm:"type:varchar(255);column:title"`
		Type                  uint64 `gorm:"type:tinyint unsigned;column:type"`
		RedirectTo            uint64 `gorm:"type:tinyint unsigned;column:redirect_to"`
		MaterialId            uint64 `gorm:"type:bigint unsigned;column:material_id"`
		MaterialType          uint64 `gorm:"type:tinyint unsigned;column:material_type"`
		Size                  string `gorm:"type:varchar(60);column:size"`
		RedirectParams        string `gorm:"type:varchar(60);column:redirect_params"`
		Description           string `gorm:"type:text;column:description"`
		Status                uint64 `gorm:"type:tinyint unsigned;column:status"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Advertising{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Advertising{})
	}

	migrate.Add("2022_04_01_170257_create_advertisings_table", up, down)
}
