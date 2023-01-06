package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
