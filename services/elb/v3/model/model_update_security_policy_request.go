package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSecurityPolicyRequest struct {
	SecurityPolicyId string `json:"security_policy_id"`

	Body *UpdateSecurityPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyRequest", string(data)}, " ")
}
