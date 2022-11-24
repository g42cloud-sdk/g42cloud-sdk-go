package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeTransferDetail struct {
	AuthKey string `json:"auth_key"`

	CreatedAt string `json:"created_at"`

	Id string `json:"id"`

	Links []Link `json:"links"`

	Name string `json:"name"`

	VolumeId string `json:"volume_id"`
}

func (o CreateVolumeTransferDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeTransferDetail struct{}"
	}

	return strings.Join([]string{"CreateVolumeTransferDetail", string(data)}, " ")
}
