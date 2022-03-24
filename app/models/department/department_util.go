package department

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (department Department) {
    database.DB.Where("id", idstr).First(&department)
    return
}

func GetBy(field, value string) (department Department) {
    database.DB.Where("? = ?", field, value).First(&department)
    return
}

func All() (departments []Department) {
    database.DB.Find(&departments)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Department{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (departments []Department, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Department{}),
        &departments,
        app.V1URL(database.TableName(&Department{})),
        perPage,
    )
    return
}