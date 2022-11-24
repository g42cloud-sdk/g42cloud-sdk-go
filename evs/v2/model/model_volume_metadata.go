package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeMetadata struct {
	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	FullClone *string `json:"full_clone,omitempty"`

	Hwpassthrough *string `json:"hw:passthrough,omitempty"`

	OrderID *string `json:"orderID,omitempty"`
}

func (o VolumeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeMetadata struct{}"
	}

	return strings.Join([]string{"VolumeMetadata", string(data)}, " ")
}
