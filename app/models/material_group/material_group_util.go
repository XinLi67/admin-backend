package material_group

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (materialGroup MaterialGroup) {
    database.DB.Where("id", idstr).First(&materialGroup)
    return
}

func GetBy(field, value string) (materialGroup MaterialGroup) {
    database.DB.Where("? = ?", field, value).First(&materialGroup)
    return
}

func All() (materialGroups []MaterialGroup) {
    database.DB.Find(&materialGroups)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(MaterialGroup{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (materialGroups []MaterialGroup, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(MaterialGroup{}),
        &materialGroups,
        app.V1URL(database.TableName(&MaterialGroup{})),
        perPage,
    )
    return
}