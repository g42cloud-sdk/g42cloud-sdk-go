package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSubNetworkInterfacesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *[]string `json:"id,omitempty"`

	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	PrivateIpAddress *[]string `json:"private_ip_address,omitempty"`

	MacAddress *[]string `json:"mac_address,omitempty"`

	VpcId *[]string `json:"vpc_id,omitempty"`

	Description *[]string `json:"description,omitempty"`

	ParentId *[]string `json:"parent_id,omitempty"`
}

func (o ListSubNetworkInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfacesRequest", string(data)}, " ")
}
