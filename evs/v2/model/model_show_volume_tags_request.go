package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVolumeTagsRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeTagsRequest", string(data)}, " ")
}
