package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type NeutronFirewallRule struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Action NeutronFirewallRuleAction `json:"action"`

	Protocol string `json:"protocol"`

	IpVersion int32 `json:"ip_version"`

	Enabled bool `json:"enabled"`

	Public bool `json:"public"`

	DestinationIpAddress string `json:"destination_ip_address"`

	DestinationPort string `json:"destination_port"`

	SourceIpAddress string `json:"source_ip_address"`

	SourcePort string `json:"source_port"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`
}

func (o NeutronFirewallRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronFirewallRule struct{}"
	}

	return strings.Join([]string{"NeutronFirewallRule", string(data)}, " ")
}

type NeutronFirewallRuleAction struct {
	value string
}

type NeutronFirewallRuleActionEnum struct {
	DENY  NeutronFirewallRuleAction
	ALLOW NeutronFirewallRuleAction
}

func GetNeutronFirewallRuleActionEnum() NeutronFirewallRuleActionEnum {
	return NeutronFirewallRuleActionEnum{
		DENY: NeutronFirewallRuleAction{
			value: "DENY",
		},
		ALLOW: NeutronFirewallRuleAction{
			value: "ALLOW",
		},
	}
}

func (c NeutronFirewallRuleAction) Value() string {
	return c.value
}

func (c NeutronFirewallRuleAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronFirewallRuleAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
