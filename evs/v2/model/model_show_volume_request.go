package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVolumeRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeRequest", string(data)}, " ")
}
