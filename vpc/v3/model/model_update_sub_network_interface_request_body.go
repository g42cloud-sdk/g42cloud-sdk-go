package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubNetworkInterfaceRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	SubNetworkInterface *UpdateSubNetworkInterfaceOption `json:"sub_network_interface"`
}

func (o UpdateSubNetworkInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubNetworkInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubNetworkInterfaceRequestBody", string(data)}, " ")
}
