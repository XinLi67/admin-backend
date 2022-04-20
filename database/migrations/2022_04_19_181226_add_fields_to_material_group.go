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

        ParentId uint64 `gorm:"type:bigint;not null;index"`
        Path     string `gorm:"type:varchar(255);not null;index"`
        

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&MaterialGroup{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&MaterialGroup{}, "ParentId")
         migrator.DropTable(&MaterialGroup{}, "Path")
    }

    migrate.Add("2022_04_19_181226_add_fields_to_material_group", up, down)
}