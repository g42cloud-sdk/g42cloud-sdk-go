package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MergeChannelsTaskInfo struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	Description *string `json:"description,omitempty"`

	AudioFiles *[]AudioFile `json:"audio_files,omitempty"`

	OutputFilename *string `json:"output_filename,omitempty"`
}

func (o MergeChannelsTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeChannelsTaskInfo struct{}"
	}

	return strings.Join([]string{"MergeChannelsTaskInfo", string(data)}, " ")
}
