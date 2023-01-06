package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSecurityGroupRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupRuleResponse", string(data)}, " ")
}
