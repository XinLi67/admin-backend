package click_record

import (
    "gohub/app/models/click_record_group_result"
    "gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (clickRecord ClickRecord) {
	database.DB.Preload("Advertising").Where("id", idstr).First(&clickRecord)
	return
}

func GetBy(field, value string) (clickRecord ClickRecord) {
	database.DB.Where("? = ?", field, value).First(&clickRecord)
	return
}

func All() (clickRecords []ClickRecord) {
	database.DB.Preload("Advertising").Find(&clickRecords)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(ClickRecord{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (clickRecords []ClickRecord, paging paginator.Paging) {
//<<<<<<< Updated upstream
	paging = paginator.Paginate(
		c,
		database.DB.Model(ClickRecord{}),
		&clickRecords,
		app.V1URL(database.TableName(&ClickRecord{})),
		perPage,
	)
	return
}
//=======
//    paging = paginator.Paginate(
//        c,
//        database.DB.Model(ClickRecord{}),
//        &clickRecords,
//        app.V1URL(database.TableName(&ClickRecord{})),
//        perPage,
//    )
//    return
//}

//按月份统计点击次数
//func ShowByMonth() (clickCount []CountRecord) {
//    database.DB.Model(&ClickRecord{}).Select("DATE_FORMAT(created_at,'%m') as mon, count(id) as total").Group("DATE_FORMAT(created_at,'%Y-%m')").Find(&clickCount)
//    return
//}

func ShowByMonth() (clickCount []DateGroupResult) {
    database.DB.Model(&ClickRecord{}).Select("DATE_FORMAT(created_at,'%m') as mon, count(id) as total").Group("DATE_FORMAT(created_at,'%Y-%m')").Find(&clickCount)
    return
}

func ShowByWeek() (clickCount []DateGroupResult) {
    database.DB.Model(&ClickRecord{}).Select("DATE_FORMAT(created_at,'%w') as week, count(id) as total").Group("DATE_FORMAT(created_at,'%Y-%w')").Find(&clickCount)
    return
}

func ShowByAdversingId() (adversingClickCount []click_record_group_result.ClickRecordGroupResult) {
    database.DB.Model(&ClickRecord{}).Select("advertising_id,count(id) as total ").Group("advertising_id").Find(&adversingClickCount)
    return
}

func ShowByCustomerId() (adversingClickCount []click_record_group_result.ClickRecordGroupResult) {
    database.DB.Model(&ClickRecord{}).Preload("admin_user").Select("customer_id,count(id) as total ").Group("customer_id").Find(&adversingClickCount)
    return
}

func ShowByPosId() (adversingClickCount []click_record_group_result.ClickRecordGroupResult) {
    //database.DB.Preload("advertising_position").Preload("User").Find(&advertisings)
    database.DB.Model(&ClickRecord{}).Preload("advertising_position").Select("position_id as advertising_position_id ,count(id) as total ").Group("position_id").Find(&adversingClickCount)
    return
}
//>>>>>>> Stashed changes
