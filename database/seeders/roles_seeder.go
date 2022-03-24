package seeders

import (
	"fmt"
	"gohub/app/models/role"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedRolesTable", func(db *gorm.DB) {

		roles := []role.Role{
			{
				GuardName:   "admin",
				Name:        "admin",
				Description: "",
			},
		}

		result := db.Table("roles").Create(&roles)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
