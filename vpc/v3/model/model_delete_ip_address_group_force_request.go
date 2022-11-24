package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteIpAddressGroupForceRequest struct {
	AddressGroupId string `json:"address_group_id"`
}

func (o DeleteIpAddressGroupForceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpAddressGroupForceRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpAddressGroupForceRequest", string(data)}, " ")
}
