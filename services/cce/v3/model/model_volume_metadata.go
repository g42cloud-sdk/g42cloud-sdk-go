package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VolumeMetadata struct {
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	SystemCmkid *string `json:"__system__cmkid,omitempty"`
}

func (o VolumeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeMetadata struct{}"
	}

	return strings.Join([]string{"VolumeMetadata", string(data)}, " ")
}
