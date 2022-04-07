package seeders

import (
	"fmt"
	"gohub/app/models/advertising_position"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAdvertisingPositionsTable", func(db *gorm.DB) {

		advertisingPositions := []advertising_position.AdvertisingPosition{
			{
				ChannelId:   1,
				Name:        "广告位一",
				Description: "这是广告位一",
				Status:      0,
				Code:        "P001",
				Height:      500,
				Weight:      700,
			},
			{
				ChannelId:   1,
				Name:        "广告位二",
				Description: "这是广告位二",
				Status:      1,
				Code:        "P002",
				Height:      500,
				Weight:      700,
			},
		}

		result := db.Table("advertising_positions").Create(&advertisingPositions)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
