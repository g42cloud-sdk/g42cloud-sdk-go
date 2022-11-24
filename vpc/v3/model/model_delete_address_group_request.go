package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAddressGroupRequest struct {
	AddressGroupId string `json:"address_group_id"`
}

func (o DeleteAddressGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddressGroupRequest", string(data)}, " ")
}
