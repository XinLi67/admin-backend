package seeders

import (
	"fmt"
	"gohub/app/models/announcement_plan"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAnnouncementPlansTable", func(db *gorm.DB) {

		announcementPlans := []announcement_plan.AnnouncementPlan{
			{
				Name:                   "公告计划一",
				CreatorId:              1,
				AnnouncementId:         1,
				AnnouncementType:       1,
				AnnouncementPositionId: 1,
				SchedulingDate:         1,
				SchedulingTime:         1,
				Order:                  1,
				AuditStatus:            0,
				PresentStatus:          0,
				StartTime:              carbon.Now().ToTimeString(),
				StartDate:              carbon.Now().ToDateString(),
				Endime:                 carbon.Tomorrow().ToTimeString(),
				EndDate:                carbon.Tomorrow().ToDateString(),
			},
			{
				Name:                   "公告计划二",
				CreatorId:              1,
				AnnouncementId:         1,
				AnnouncementType:       1,
				AnnouncementPositionId: 1,
				SchedulingDate:         1,
				SchedulingTime:         1,
				Order:                  2,
				AuditStatus:            1,
				PresentStatus:          1,
				StartTime:              carbon.Now().ToTimeString(),
				StartDate:              carbon.Now().ToDateString(),
				Endime:                 carbon.Tomorrow().ToTimeString(),
				EndDate:                carbon.Tomorrow().ToDateString(),
			},
		}

		result := db.Table("announcement_plans").Create(&announcementPlans)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
