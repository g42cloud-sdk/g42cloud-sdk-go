package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallPolicyRequestBody struct {
	FirewallPolicy *NeutronCreateFirewallPolicyOption `json:"firewall_policy"`
}

func (o NeutronCreateFirewallPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallPolicyRequestBody", string(data)}, " ")
}
