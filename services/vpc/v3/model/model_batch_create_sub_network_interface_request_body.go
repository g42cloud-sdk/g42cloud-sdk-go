package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
