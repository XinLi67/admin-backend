package policies

import (
	"gohub/app/models/audit_record"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
// func CanCreateAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
// func CanUpdateAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
// func CanDeleteAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
