package seeders

import (
	"fmt"
	"gohub/app/models/announcement"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAnnouncementsTable", func(db *gorm.DB) {

		announcements := []announcement.Announcement{
			{
				AnnouncementNo:         202201010000,
				AnnouncementPositionId: 1,
				CreatorId:              1,
				DepartmentId:           1,
				Title:                  "公告一",
				Type:                   0,
				Status:                 1,
			},
			{
				AnnouncementNo:         202201010000,
				AnnouncementPositionId: 1,
				CreatorId:              1,
				DepartmentId:           1,
				Title:                  "公告二",
				Type:                   1,
				Status:                 1,
			},
		}

		result := db.Table("announcements").Create(&announcements)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
