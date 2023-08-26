package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type NeutronCreateFirewallRuleOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	Action *NeutronCreateFirewallRuleOptionAction `json:"action,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`

	DestinationIpAddress *string `json:"destination_ip_address,omitempty"`

	DestinationPort *string `json:"destination_port,omitempty"`

	SourceIpAddress *string `json:"source_ip_address,omitempty"`

	SourcePort *string `json:"source_port,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

func (o NeutronCreateFirewallRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleOption", string(data)}, " ")
}

type NeutronCreateFirewallRuleOptionAction struct {
	value string
}

type NeutronCreateFirewallRuleOptionActionEnum struct {
	DENY  NeutronCreateFirewallRuleOptionAction
	ALLOW NeutronCreateFirewallRuleOptionAction
}

func GetNeutronCreateFirewallRuleOptionActionEnum() NeutronCreateFirewallRuleOptionActionEnum {
	return NeutronCreateFirewallRuleOptionActionEnum{
		DENY: NeutronCreateFirewallRuleOptionAction{
			value: "DENY",
		},
		ALLOW: NeutronCreateFirewallRuleOptionAction{
			value: "ALLOW",
		},
	}
}

func (c NeutronCreateFirewallRuleOptionAction) Value() string {
	return c.value
}

func (c NeutronCreateFirewallRuleOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronCreateFirewallRuleOptionAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
