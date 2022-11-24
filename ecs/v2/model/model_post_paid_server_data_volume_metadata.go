package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostPaidServerDataVolumeMetadata struct {
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	SystemCmkid *string `json:"__system__cmkid,omitempty"`
}

func (o PostPaidServerDataVolumeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerDataVolumeMetadata struct{}"
	}

	return strings.Join([]string{"PostPaidServerDataVolumeMetadata", string(data)}, " ")
}
