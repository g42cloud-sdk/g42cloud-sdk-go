package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateSecurityGroupRuleResponse struct {
	SecurityGroupRule *NeutronSecurityGroupRule `json:"security_group_rule,omitempty"`
	HttpStatusCode    int                       `json:"-"`
}

func (o NeutronCreateSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleResponse", string(data)}, " ")
}
