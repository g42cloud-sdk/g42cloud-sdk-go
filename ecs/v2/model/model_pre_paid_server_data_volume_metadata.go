package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrePaidServerDataVolumeMetadata struct {
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	SystemCmkid *string `json:"__system__cmkid,omitempty"`
}

func (o PrePaidServerDataVolumeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerDataVolumeMetadata struct{}"
	}

	return strings.Join([]string{"PrePaidServerDataVolumeMetadata", string(data)}, " ")
}
