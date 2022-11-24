package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubNetworkInterfaceRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	SubNetworkInterface *CreateSubNetworkInterfaceOption `json:"sub_network_interface"`
}

func (o CreateSubNetworkInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceRequestBody", string(data)}, " ")
}
