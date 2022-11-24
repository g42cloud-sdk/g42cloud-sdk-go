package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAuditlogPolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowAuditlogPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditlogPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditlogPolicyRequest", string(data)}, " ")
}
