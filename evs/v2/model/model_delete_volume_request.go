package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteVolumeRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o DeleteVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeRequest struct{}"
	}

	return strings.Join([]string{"DeleteVolumeRequest", string(data)}, " ")
}
