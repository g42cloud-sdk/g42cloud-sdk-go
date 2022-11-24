package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeSchedulerHints struct {
	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`
}

func (o CreateVolumeSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeSchedulerHints struct{}"
	}

	return strings.Join([]string{"CreateVolumeSchedulerHints", string(data)}, " ")
}
