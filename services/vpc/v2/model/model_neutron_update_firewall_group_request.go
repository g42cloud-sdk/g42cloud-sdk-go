package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallGroupRequest struct {
	FirewallGroupId string `json:"firewall_group_id"`

	Body *NeutronUpdateFirewallGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateFirewallGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupRequest", string(data)}, " ")
}
