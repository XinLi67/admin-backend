package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AuditRecordRequest struct {
	AuditableId   uint64 `json:"auditable_id,omitempty"`
	AuditableType uint64 `valid:"auditable_type" json:"auditable_type,omitempty"`
	ApplicantId   uint64 `json:"applicant_id,omitempty"`
	AuditorId     uint64 `json:"auditor_id,omitempty"`
	Status        uint64 `json:"status,omitempty"`
	Content       string `json:"content,omitempty"`
}

func AuditRecordSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"auditable_type": []string{"in:0,1,2,3"},
	}
	messages := govalidator.MapData{
		"auditable_type": []string{
			"in:只能为0或1或2或3",
		},
	}
	return validate(data, rules, messages)
}
