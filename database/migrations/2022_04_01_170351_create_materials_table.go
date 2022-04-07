package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Material struct {
		models.BaseModel

		CreatorId       uint64 `gorm:"type:bigint unsigned;column:creator_id"`
		MaterialGroupId uint64 `gorm:"type:bigint unsigned;column:material_group_id"`
		DepartmentId    uint64 `gorm:"type:bigint unsigned;column:department_id"`
		Type            uint64 `gorm:"type:tinyint unsigned;column:type"`
		Url             string `gorm:"type:varchar(255);column:url"`
		Title           string `gorm:"type:varchar(255);column:title"`
		Content         string `gorm:"type:text;column:content"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Material{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Material{})
	}

	migrate.Add("2022_04_01_170351_create_materials_table", up, down)
}
