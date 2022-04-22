package advertising

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idstr string) (advertising Advertising) {
	database.DB.Preload("AdvertisingPosition").Preload("User").Where("id", idstr).First(&advertising)
	return
}

//根据广告位获取对应广告
func GetAll(idstr string) (advertising []Advertising) {
	// database.DB.Where("id", idstr).First(&advertising)
	database.DB.Preload("advertising_position").Where("advertising_position_id", idstr).Find(&advertising)
	return
}

func GetBy(field, value string) (advertising Advertising) {
	database.DB.Where("? = ?", field, value).First(&advertising)
	return
}

func All() (advertisings []Advertising) {
	// database.DB.Find(&advertisings)
	database.DB.Preload("advertising_position").Preload("User").Find(&advertisings)
	return
}

func All2() (advertisings []*Advertising) {
	// database.DB.Find(&advertisings)
	database.DB.Find(&advertisings)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Advertising{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})),
		perPage,
	)
	return
}

func PaginateByTitle(c *gin.Context, perPage int, params string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		//database.DB.Model(Material{}),
		database.DB.Model(Advertising{}).Where("id like ?", "%"+params+"%").
			//Or("creator_id like ?", "%"+params+"%").
			//Or("advertising_no like ?", "%"+params+"%").
			//Or("department_id like ?", "%"+params+"%").
			Or("title like ?", "%"+params+"%"),
		//Or("type like ?", "%"+params+"%").
		//Or("material_id like ?", "%"+params+"%").
		//Or("material_type like ?", "%"+params+"%").
		//Or("size like ?", "%"+params+"%").
		//Or("redirect_to like ?", "%"+params+"%").
		//Or("redirect_params like ?", "%"+params+"%").
		//Or("description like ?", "%"+params+"%"),

		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?params="+params),
		perPage,
	)
	return advertisings, paging
}

func PaginateByType(c *gin.Context, perPage int, adtype string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("type = ?", adtype),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?type="+adtype),
		perPage,
	)
	return advertisings, paging
}

func PaginateByPosId(c *gin.Context, perPage int, posId string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("advertising_position_id = ?", posId),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?advertising_position_id="+posId),
		perPage,
	)
	return advertisings, paging
}

func PaginateByAdtypeAndAdvertisingPositionId(c *gin.Context, perPage int, adtype string,posId string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("type = ? and advertising_position_id=?", adtype,posId),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?type="+adtype+"?advertising_position_id="+posId),
		perPage,
	)
	return advertisings, paging
}

func PaginateByAdtypeAndParams(c *gin.Context, perPage int, adtype string,params string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("type = ? and title like ?", adtype,"%"+params+"%"),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?type="+adtype+"?params="+params),
		perPage,
	)
	return advertisings, paging
}

//根据审核状态查询后分页显示
func PaginateByStatus(c *gin.Context, perPage int, status string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("status = ?", status),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?status="+status),
		perPage,
	)
	return advertisings, paging
}

//根据审核状态和模糊查询参数名查询后分页显示
func PaginateByStatusAndParams(c *gin.Context, perPage int, status string, params string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("status = ? and title like ?", status, "%"+params+"%"),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?status="+status+"?params="+params),
		perPage,
	)
	return advertisings, paging
}

func PaginateByParamsAndAdtypeAndAdvertisingPositionId(c *gin.Context, perPage int,  params string, adtype string,posId string) (advertisings []Advertising, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Advertising{}).Where("title like ? and type =  ? and advertising_position_id=?", "%"+params+"%", adtype,posId),
		&advertisings,
		app.V1URL(database.TableName(&Advertising{})+"/list?params="+params+"?type="+adtype+"advertisingPositionId="+posId),
		perPage,
	)
	return advertisings, paging
}
