package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowServerBlockDeviceResponse struct {
	VolumeAttachment *ServerBlockDevice `json:"volumeAttachment,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ShowServerBlockDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerBlockDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowServerBlockDeviceResponse", string(data)}, " ")
}
