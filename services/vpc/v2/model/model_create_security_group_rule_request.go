package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRuleRequest struct {
	Body *CreateSecurityGroupRuleRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleRequest", string(data)}, " ")
}
