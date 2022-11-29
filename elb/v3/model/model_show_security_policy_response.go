package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSecurityPolicyResponse struct {
	SecurityPolicy *SecurityPolicy `json:"security_policy,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityPolicyResponse", string(data)}, " ")
}
