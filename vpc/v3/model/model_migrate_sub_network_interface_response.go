package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateSubNetworkInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterfaces *[]SubNetworkInterface `json:"sub_network_interfaces,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o MigrateSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"MigrateSubNetworkInterfaceResponse", string(data)}, " ")
}
