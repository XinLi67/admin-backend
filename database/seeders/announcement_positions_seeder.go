package seeders

import (
	"fmt"
	"gohub/app/models/announcement_position"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAnnouncementPositionsTable", func(db *gorm.DB) {

		announcementPositions := []announcement_position.AnnouncementPosition{
			{
				ChannelId:   1,
				Name:        "公告位一",
				Description: "这是公告位一",
				Status:      0,
				Code:        "P001",
				Height:      100,
				Weight:      200,
			},
			{
				ChannelId:   1,
				Name:        "公告位二",
				Description: "这是公告位二",
				Status:      1,
				Code:        "P002",
				Height:      100,
				Weight:      200,
			},
		}

		result := db.Table("announcement_positions").Create(&announcementPositions)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
