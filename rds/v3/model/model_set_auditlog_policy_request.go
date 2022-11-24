package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAuditlogPolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SetAuditlogPolicyRequestBody `json:"body,omitempty"`
}

func (o SetAuditlogPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyRequest", string(data)}, " ")
}
