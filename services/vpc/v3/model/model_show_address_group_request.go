package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowAddressGroupRequest struct {
	AddressGroupId string `json:"address_group_id"`
}

func (o ShowAddressGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddressGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowAddressGroupRequest", string(data)}, " ")
}
