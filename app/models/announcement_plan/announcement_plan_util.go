package announcement_plan

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (announcementPlan AnnouncementPlan) {
	// database.DB.Where("id", idstr).First(&announcementPlan)
	// database.DB.Preload("User").Preload("AnnouncementPosition").Where("id", idstr).First(&announcementPlan)
	database.DB.Preload("User").Where("id", idstr).First(&announcementPlan)
	return
}

func GetBy(field, value string) (announcementPlan AnnouncementPlan) {
	database.DB.Where("? = ?", field, value).First(&announcementPlan)
	return
}

func All() (announcementPlans []AnnouncementPlan) {
	// database.DB.Find(&announcementPlans)
	database.DB.Preload("User").Preload("AnnouncementPosition").Find(&announcementPlans)
	return
}

func All2() (announcementPlans []*AnnouncementPlan) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&announcementPlans)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AnnouncementPlan{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPlan{}),
		&announcementPlans,
		app.V1URL(database.TableName(&AnnouncementPlan{})),
		perPage,
	)
	return
}

func PaginateByName(c *gin.Context, perPage int, params string) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPlan{}).Where("name like ?", "%"+params+"%"),
		&announcementPlans,
		app.V1URL(database.TableName(&AnnouncementPlan{})+"/list?params="+params),
		perPage,
	)
	return announcementPlans, paging
}

//根据审核状态查询后分页显示
func PaginateByStatus(c *gin.Context, perPage int, audit_status string) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPlan{}).Where("audit_status = ?", audit_status),
		&announcementPlans,
		app.V1URL(database.TableName(&AnnouncementPlan{})+"/list?audit_status="+audit_status),
		perPage,
	)
	return announcementPlans, paging
}

//根据审核状态查询后分页显示
func PaginateByStatusAndParams(c *gin.Context, perPage int, audit_status string, params string) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(AnnouncementPlan{}).Where("audit_status = ? and name like ?", audit_status, "%"+params+"%"),
		&announcementPlans,
		app.V1URL(database.TableName(&AnnouncementPlan{})+"/list?status="+audit_status+"?params="+params),
		perPage,
	)
	return announcementPlans, paging
}
