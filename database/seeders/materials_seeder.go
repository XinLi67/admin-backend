package seeders

import (
	"fmt"
	"gohub/app/models/material"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedMaterialsTable", func(db *gorm.DB) {

		materials := []material.Material{
			{
				CreatorId:       1,
				MaterialGroupId: 1,
				DepartmentId:    1,
				Type:            0,
				Title:           "素材一",
				Content:         "这是素材一",
			},
			{
				CreatorId:       1,
				MaterialGroupId: 2,
				DepartmentId:    1,
				Type:            0,
				Title:           "素材二",
				Content:         "这是素材二",
			},
			{
				CreatorId:       1,
				MaterialGroupId: 3,
				DepartmentId:    1,
				Type:            0,
				Title:           "素材三",
				Content:         "这是素材三",
			},
		}

		result := db.Table("materials").Create(&materials)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
