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

    seed.Add("SeedDepartmentsTable", func(db *gorm.DB) {

        departments  := factories.MakeDepartments(10)

        result := db.Table("departments").Create(&departments)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}