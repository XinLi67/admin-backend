package seeders

import (
	"fmt"
	"gohub/app/models/channel"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedChannelsTable", func(db *gorm.DB) {

		channels := []channel.Channel{
			{
				Name:        "手机银行",
				Description: "这是手机银行",
				Status:      1,
			},
			{
				Name:        "微信银行",
				Description: "这是微信银行",
				Status:      1,
			},
		}

		result := db.Table("channels").Create(&channels)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
