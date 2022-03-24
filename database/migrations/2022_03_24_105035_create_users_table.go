package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		DepartmentId uint64 `gorm:"type:bigint unsigned;not null;index;comment:部门ID"`
		GuardName    string `gorm:"type:varchar(30);not null;index;comment:项目"`
		Username     string `gorm:"type:varchar(60);not null;comment:登录名称"`
		Name         string `gorm:"type:varchar(60);not null;comment:用户名称"`
		Gender       uint64 `gorm:"type:tinyint unsigned;default:1;comment:性别 1-男 2-女"`
		Email        string `gorm:"type:varchar(60);index;default:null;comment:邮箱"`
		Phone        string `gorm:"type:varchar(20);index;default:null;comment:电话"`
		Avatar       string `gorm:"type:varchar(255);default:null;comment:头象"`
		Password     string `gorm:"type:varchar(255);comment:密码"`
		Status       uint64 `gorm:"tinyint unsigned;not null;default 0;comment:生效状态"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_03_24_105035_create_users_table", up, down)
}
