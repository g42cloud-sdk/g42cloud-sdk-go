package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MediaProcessTaskInfo struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	Description *string `json:"description,omitempty"`

	OutputFileName *[]string `json:"output_file_name,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`
}

func (o MediaProcessTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaProcessTaskInfo struct{}"
	}

	return strings.Join([]string{"MediaProcessTaskInfo", string(data)}, " ")
}
