package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
