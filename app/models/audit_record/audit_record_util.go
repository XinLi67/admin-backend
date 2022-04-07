package audit_record

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (auditRecord AuditRecord) {
    database.DB.Where("id", idstr).First(&auditRecord)
    return
}

func GetBy(field, value string) (auditRecord AuditRecord) {
    database.DB.Where("? = ?", field, value).First(&auditRecord)
    return
}

func All() (auditRecords []AuditRecord) {
    database.DB.Find(&auditRecords)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(AuditRecord{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (auditRecords []AuditRecord, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(AuditRecord{}),
        &auditRecords,
        app.V1URL(database.TableName(&AuditRecord{})),
        perPage,
    )
    return
}