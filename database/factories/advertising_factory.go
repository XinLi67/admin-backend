package factories

import (
	"github.com/golang-module/carbon/v2"
	"gohub/app/models/advertising"
	"gohub/pkg/helpers"
)

func MakeAdvertisings(count int) []advertising.Advertising {

	var objs []advertising.Advertising

	// 设置唯一性，如 Advertising 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 1; i <= count; i++ {
		advertisingModel := advertising.Advertising{
			AdvertisingNo:         uint64(220414115732549211 + i),
			AdvertisingPositionId: 1,
			CreatorId:             uint64(i + 1),
			Title:                 helpers.RandomString(5),
			Type:                  1,
			MaterialId:            1,
			MaterialType:          0,
			Size:                  "800x600",
			RedirectTo:            1,
			DepartmentId:          1,
			RedirectParams:        helpers.RandomString(5),
			Description:           helpers.RandomString(5),
			StartTime:             carbon.Now().ToTimeString(),
			EndTime:               carbon.Tomorrow().ToTimeString(),
		}
		objs = append(objs, advertisingModel)
	}

	return objs
}
