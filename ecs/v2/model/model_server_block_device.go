package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerBlockDevice struct {
	BootIndex *int32 `json:"bootIndex,omitempty"`

	PciAddress *string `json:"pciAddress,omitempty"`

	VolumeId *string `json:"volumeId,omitempty"`

	Device *string `json:"device,omitempty"`

	ServerId *string `json:"serverId,omitempty"`

	Id *string `json:"id,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Bus *string `json:"bus,omitempty"`
}

func (o ServerBlockDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerBlockDevice struct{}"
	}

	return strings.Join([]string{"ServerBlockDevice", string(data)}, " ")
}
