package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Department struct {
		models.BaseModel

		ParentId    uint64 `gorm:"type:bigint unsigned;not null;index;default:0;comment:上级ID"`
		Name        string `gorm:"type:varchar(120);not null;comment:部门名称"`
		Phone       string `gorm:"type:varchar(60);null;comment:部门电话"`
		LinkMan     string `gorm:"type:varchar(60);null;comment:部门联系人"`
		Address     string `gorm:"type:varchar(255);null;comment:部门地址"`
		Description string `gorm:"type:varchar(255);null;comment:部门简介"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Department{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Department{})
	}

	migrate.Add("2022_03_24_105015_create_departments_table", up, down)
}
