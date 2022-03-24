package seeders

import (
	"fmt"
	"gohub/app/models/department"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedDepartmentsTable", func(db *gorm.DB) {

		departments := []department.Department{
			{
				ParentId:    0,
				Name:        "总行",
				Phone:       "18007729523",
				LinkMan:     "雷德钦",
				Description: "这个是总行部门",
				Address:     "柳州市西堤路12号",
			}, {
				ParentId:    1,
				Name:        "数字银行部",
				Phone:       "18007729523",
				LinkMan:     "雷德钦",
				Description: "这个是数字银行部",
				Address:     "柳州市西堤路12号",
			}, {
				ParentId:    1,
				Name:        "信息科技部",
				Phone:       "18007729523",
				LinkMan:     "雷德钦",
				Description: "这个是信息科技部",
				Address:     "柳州市西堤路12号",
			}, {
				ParentId:    1,
				Name:        "公司业务部",
				Phone:       "18007729523",
				LinkMan:     "雷德钦",
				Description: "这个是公司业务部",
				Address:     "柳州市西堤路12号",
			}, {
				ParentId:    1,
				Name:        "风险管理部",
				Phone:       "18007729523",
				LinkMan:     "雷德钦",
				Description: "这个是风险管理部",
				Address:     "柳州市西堤路12号",
			}, {
				ParentId:    1,
				Name:        "个人金融部",
				Phone:       "18007729523",
				LinkMan:     "雷德钦",
				Description: "这个是个人金融部",
				Address:     "柳州市西堤路12号",
			},
		}

		result := db.Table("departments").Create(&departments)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
