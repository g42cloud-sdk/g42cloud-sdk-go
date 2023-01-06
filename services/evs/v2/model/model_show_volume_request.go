package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
