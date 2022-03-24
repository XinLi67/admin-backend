package seeders

import (
	"fmt"
	"gohub/app/models/user"
	"gohub/pkg/console"
	"gohub/pkg/hash"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := []user.User{
			{
				DepartmentId: 1,
				GuardName:    "admin",
				Username:     "admin",
				Name:         "admin",
				Gender:       1,
				Email:        "leideqin@126.com",
				Phone:        "18207723661",
				Avatar:       "http://dummyimage.com/100x100",
				Password:     hash.BcryptHash("secret"),
				Status:       0,
			}, {
				DepartmentId: 2,
				GuardName:    "admin",
				Username:     "rebud",
				Name:         "rebud",
				Gender:       1,
				Email:        "leideqin@aliyun.com",
				Phone:        "18007729523",
				Avatar:       "http://dummyimage.com/100x100",
				Password:     hash.BcryptHash("secret"),
				Status:       0,
			},
		}

		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
