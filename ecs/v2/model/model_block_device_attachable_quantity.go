package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BlockDeviceAttachableQuantity struct {
	FreeScsi *int32 `json:"free_scsi,omitempty"`

	FreeBlk *int32 `json:"free_blk,omitempty"`

	FreeDisk *int32 `json:"free_disk,omitempty"`
}

func (o BlockDeviceAttachableQuantity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockDeviceAttachableQuantity struct{}"
	}

	return strings.Join([]string{"BlockDeviceAttachableQuantity", string(data)}, " ")
}
