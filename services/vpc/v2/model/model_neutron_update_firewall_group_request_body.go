package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronUpdateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronUpdateFirewallGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupRequestBody", string(data)}, " ")
}
