package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteFirewallPolicyRequest struct {
	FirewallPolicyId string `json:"firewall_policy_id"`
}

func (o NeutronDeleteFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallPolicyRequest", string(data)}, " ")
}
