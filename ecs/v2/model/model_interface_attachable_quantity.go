package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterfaceAttachableQuantity struct {
	FreeNic *int32 `json:"free_nic,omitempty"`
}

func (o InterfaceAttachableQuantity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfaceAttachableQuantity struct{}"
	}

	return strings.Join([]string{"InterfaceAttachableQuantity", string(data)}, " ")
}
