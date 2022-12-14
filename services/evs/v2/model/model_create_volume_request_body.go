package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVolumeRequestBody struct {
	BssParam *BssParamForCreateVolume `json:"bssParam,omitempty"`

	Volume *CreateVolumeOption `json:"volume"`

	ServerId *string `json:"server_id,omitempty"`

	OSSCHHNTschedulerHints *CreateVolumeSchedulerHints `json:"OS-SCH-HNT:scheduler_hints,omitempty"`
}

func (o CreateVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVolumeRequestBody", string(data)}, " ")
}
