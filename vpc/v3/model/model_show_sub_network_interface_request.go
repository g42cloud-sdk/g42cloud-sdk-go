package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSubNetworkInterfaceRequest struct {
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`
}

func (o ShowSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfaceRequest", string(data)}, " ")
}
