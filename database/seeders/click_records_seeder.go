package seeders

import (
	"fmt"
	"gohub/app/models/click_record"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedClickRecordsTable", func(db *gorm.DB) {

		clickRecords := []click_record.ClickRecord{
			{
				AdvertisingId: 1,
				CustomerId:    1,
				BrowsingTime:  60,
			},
			{
				AdvertisingId: 1,
				CustomerId:    2,
				BrowsingTime:  60,
			},
			{
				AdvertisingId: 1,
				CustomerId:    3,
				BrowsingTime:  60,
			},
			{
				AdvertisingId: 2,
				CustomerId:    1,
				BrowsingTime:  60,
			},
			{
				AdvertisingId: 2,
				CustomerId:    2,
				BrowsingTime:  60,
			},
			{
				AdvertisingId: 2,
				CustomerId:    3,
				BrowsingTime:  60,
			},
		}

		result := db.Table("click_records").Create(&clickRecords)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
