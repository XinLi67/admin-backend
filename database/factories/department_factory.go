package factories

import (
	"gohub/app/models/department"

	"github.com/bxcodec/faker/v3"
)

func MakeDepartments(count int) []department.Department {

	var objs []department.Department

	// 设置唯一性，如 Department 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	parentId, _ := faker.RandomInt(1, 2)
	println(parentId)

	for i := 0; i < count; i++ {
		departmentModel := department.Department{
			ParentId:    uint64(parentId[0]),
			Name:        faker.Name(),
			Phone:       faker.Phonenumber(),
			LinkMan:     faker.FirstName(),
			Description: faker.Sentence(),
			Address:     faker.Sentence(),
		}
		objs = append(objs, departmentModel)
	}

	return objs
}
