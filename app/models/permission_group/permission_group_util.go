package permission_group

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (permissionGroup PermissionGroup) {
	database.DB.Where("id", idstr).First(&permissionGroup)
	return
}

func GetBy(field, value string) (permissionGroup PermissionGroup) {
	database.DB.Where("? = ?", field, value).First(&permissionGroup)
	return
}

func GetByName(name string) (permissionGroup PermissionGroup) {
	database.DB.Where("name = ?", name).First(&permissionGroup)
	return
}

func All() (permissionGroups []PermissionGroup) {
	database.DB.Find(&permissionGroups)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PermissionGroup{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (permissionGroups []PermissionGroup, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(PermissionGroup{}),
		&permissionGroups,
		app.V1URL(database.TableName(&PermissionGroup{})),
		perPage,
	)
	return
}
