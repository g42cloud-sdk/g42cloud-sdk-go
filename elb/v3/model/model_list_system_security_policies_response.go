package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSystemSecurityPoliciesResponse struct {
	SystemSecurityPolicies *[]SystemSecurityPolicy `json:"system_security_policies,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSystemSecurityPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemSecurityPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSystemSecurityPoliciesResponse", string(data)}, " ")
}
