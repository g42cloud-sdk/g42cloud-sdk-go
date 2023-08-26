package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSecurityGroupRuleResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ShowSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupRuleResponse", string(data)}, " ")
}
