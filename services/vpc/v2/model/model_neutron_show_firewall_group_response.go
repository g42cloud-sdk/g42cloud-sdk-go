package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronShowFirewallGroupResponse struct {
	FirewallGroup  *NeutronFirewallGroup `json:"firewall_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronShowFirewallGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallGroupResponse", string(data)}, " ")
}
