package click_record_group_result

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (clickRecordGroupResult ClickRecordGroupResult) {
    database.DB.Where("id", idstr).First(&clickRecordGroupResult)
    return
}

func GetBy(field, value string) (clickRecordGroupResult ClickRecordGroupResult) {
    database.DB.Where("? = ?", field, value).First(&clickRecordGroupResult)
    return
}

func All() (clickRecordGroupResults []ClickRecordGroupResult) {
    database.DB.Find(&clickRecordGroupResults)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(ClickRecordGroupResult{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (clickRecordGroupResults []ClickRecordGroupResult, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(ClickRecordGroupResult{}),
        &clickRecordGroupResults,
        app.V1URL(database.TableName(&ClickRecordGroupResult{})),
        perPage,
    )
    return
}