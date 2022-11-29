package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSecurityPolicyResponse struct {
	SecurityPolicy *SecurityPolicy `json:"security_policy,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyResponse", string(data)}, " ")
}