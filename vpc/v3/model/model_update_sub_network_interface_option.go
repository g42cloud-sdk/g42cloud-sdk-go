package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubNetworkInterfaceOption struct {
	Description *string `json:"description,omitempty"`

	SecurityGroups *[]string `json:"security_groups,omitempty"`
}

func (o UpdateSubNetworkInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubNetworkInterfaceOption struct{}"
	}

	return strings.Join([]string{"UpdateSubNetworkInterfaceOption", string(data)}, " ")
}
