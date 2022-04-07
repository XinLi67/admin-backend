package seeders

import (
	"fmt"
	"gohub/app/models/advertising"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAdvertisingsTable", func(db *gorm.DB) {

		advertisings := []advertising.Advertising{
			{
				AdvertisingNo:         202201010000,
				AdvertisingPositionId: 1,
				CreatorId:             1,
				DepartmentId:          1,
				Title:                 "广告一",
				Type:                  1,
				MaterialId:            1,
				MaterialType:          1,
				Status:                1,
			},
			{
				AdvertisingNo:         202201010001,
				AdvertisingPositionId: 1,
				CreatorId:             1,
				DepartmentId:          1,
				Title:                 "广告二",
				Type:                  2,
				MaterialId:            2,
				MaterialType:          1,
				Status:                0,
			},
		}

		result := db.Table("advertisings").Create(&advertisings)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
