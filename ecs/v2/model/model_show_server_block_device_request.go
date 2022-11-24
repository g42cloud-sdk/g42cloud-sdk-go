package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowServerBlockDeviceRequest struct {
	ServerId string `json:"server_id"`

	VolumeId string `json:"volume_id"`
}

func (o ShowServerBlockDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerBlockDeviceRequest struct{}"
	}

	return strings.Join([]string{"ShowServerBlockDeviceRequest", string(data)}, " ")
}
