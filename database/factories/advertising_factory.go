package factories

import (
	"adcenter/app/models/advertising"
	"adcenter/pkg/helpers"
)

func MakeAdvertisings(count int) []advertising.Advertising {

	var objs []advertising.Advertising

	// 设置唯一性，如 Advertising 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		advertisingModel := advertising.Advertising{
			CreatorId:      uint64(i + 1),
			Title:          helpers.RandomString(5),
			Type:           1,
			MediaId:        i * 10,
			MediaType:      1,
			Size:           "800x600",
			RedirectTo:     1,
			DepartmentId:   1,
			RedirectParams: helpers.RandomString(5),
			Description:    helpers.RandomString(5),
		}
		objs = append(objs, advertisingModel)
	}

	return objs
}
