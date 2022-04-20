package seeders

import (
	"fmt"
	"gohub/database/factories"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAdvertisingsTable", func(db *gorm.DB) {

		advertisings:=factories.MakeAdvertisings(10)

		result := db.Table("advertisings").Create(&advertisings)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
