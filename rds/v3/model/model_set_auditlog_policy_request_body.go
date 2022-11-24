package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAuditlogPolicyRequestBody struct {
	KeepDays int32 `json:"keep_days"`

	ReserveAuditlogs *bool `json:"reserve_auditlogs,omitempty"`
}

func (o SetAuditlogPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyRequestBody", string(data)}, " ")
}
