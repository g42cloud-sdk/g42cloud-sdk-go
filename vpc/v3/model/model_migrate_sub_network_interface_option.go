package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateSubNetworkInterfaceOption struct {
	ParentId string `json:"parent_id"`

	SubNetworkInterfaces []map[string]string `json:"sub_network_interfaces"`
}

func (o MigrateSubNetworkInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateSubNetworkInterfaceOption struct{}"
	}

	return strings.Join([]string{"MigrateSubNetworkInterfaceOption", string(data)}, " ")
}
