package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeTransferOption struct {
	Name string `json:"name"`

	VolumeId string `json:"volume_id"`
}

func (o CreateVolumeTransferOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeTransferOption struct{}"
	}

	return strings.Join([]string{"CreateVolumeTransferOption", string(data)}, " ")
}
