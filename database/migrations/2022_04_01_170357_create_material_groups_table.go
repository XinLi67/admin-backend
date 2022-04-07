package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type MaterialGroup struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(60);column:name"`
		Description string `gorm:"type:text;column:description"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&MaterialGroup{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&MaterialGroup{})
	}

	migrate.Add("2022_04_01_170357_create_material_groups_table", up, down)
}
