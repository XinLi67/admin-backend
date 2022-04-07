package seeders

import (
	"fmt"
	"gohub/app/models/material_group"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedMaterialGroupsTable", func(db *gorm.DB) {

		materialGroups := []material_group.MaterialGroup{
			{
				Name:        "素材组一",
				Description: "这是素材组一",
			},
			{
				Name:        "素材组二",
				Description: "这是素材组二",
			},
			{
				Name:        "素材组三",
				Description: "这是素材组三",
			},
		}

		result := db.Table("material_groups").Create(&materialGroups)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
