package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EngineReference struct {
	Vpc *string `json:"vpc,omitempty"`

	AzList *[]string `json:"az_list,omitempty"`

	NetworkId *string `json:"network_id,omitempty"`

	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	SubnetCidrV6 *string `json:"subnet_cidr_v6,omitempty"`

	SubnetGateway *string `json:"subnet_gateway,omitempty"`

	PublicIpId *string `json:"public_ip_id,omitempty"`

	ServiceLimit *int32 `json:"service_limit,omitempty"`

	InstanceLimit *int32 `json:"instance_limit,omitempty"`

	Inputs map[string]string `json:"inputs,omitempty"`
}

func (o EngineReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineReference struct{}"
	}

	return strings.Join([]string{"EngineReference", string(data)}, " ")
}
