package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSubNetworkInterfaceRequest struct {
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`
}

func (o DeleteSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubNetworkInterfaceRequest", string(data)}, " ")
}
