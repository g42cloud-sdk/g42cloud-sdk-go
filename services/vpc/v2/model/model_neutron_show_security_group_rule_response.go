package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronShowSecurityGroupRuleResponse struct {
	SecurityGroupRule *NeutronSecurityGroupRule `json:"security_group_rule,omitempty"`
	HttpStatusCode    int                       `json:"-"`
}

func (o NeutronShowSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupRuleResponse", string(data)}, " ")
}
