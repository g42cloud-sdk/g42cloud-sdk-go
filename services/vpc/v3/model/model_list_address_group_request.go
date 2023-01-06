package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAddressGroupRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`

	Description *[]string `json:"description,omitempty"`
}

func (o ListAddressGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressGroupRequest struct{}"
	}

	return strings.Join([]string{"ListAddressGroupRequest", string(data)}, " ")
}
