package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSecurityPolicyRequest struct {
	SecurityPolicyId string `json:"security_policy_id"`
}

func (o ShowSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityPolicyRequest", string(data)}, " ")
}
