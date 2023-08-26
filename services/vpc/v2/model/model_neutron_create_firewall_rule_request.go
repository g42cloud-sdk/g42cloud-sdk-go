package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallRuleRequest struct {
	Body *NeutronCreateFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleRequest", string(data)}, " ")
}
