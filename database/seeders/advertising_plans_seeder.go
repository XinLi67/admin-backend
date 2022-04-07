package seeders

import (
	"fmt"
	"gohub/app/models/advertising_plan"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAdvertisingPlansTable", func(db *gorm.DB) {

		advertisingPlans := []advertising_plan.AdvertisingPlan{
			{
				Name:                  "广告计划一",
				CreatorId:             1,
				AdvertisingId:         1,
				AdvertisingType:       1,
				AdvertisingPositionId: 1,
				SchedulingDate:        1,
				SchedulingTime:        1,
				Order:                 1,
				AuditStatus:           0,
				PresentStatus:         0,
			},
			{
				Name:                  "广告计划二",
				CreatorId:             1,
				AdvertisingId:         2,
				AdvertisingType:       1,
				AdvertisingPositionId: 2,
				SchedulingDate:        1,
				SchedulingTime:        1,
				Order:                 2,
				AuditStatus:           1,
				PresentStatus:         1,
			},
		}

		result := db.Table("advertising_plans").Create(&advertisingPlans)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
