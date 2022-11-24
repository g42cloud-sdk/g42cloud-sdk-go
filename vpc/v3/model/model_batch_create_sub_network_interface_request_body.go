package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateSubNetworkInterfaceRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	SubNetworkInterface *BatchCreateSubNetworkInterfaceOption `json:"sub_network_interface"`

	Count int32 `json:"count"`
}

func (o BatchCreateSubNetworkInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceRequestBody", string(data)}, " ")
}
