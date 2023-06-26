package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AnimatedGraphicsTask struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	Description *string `json:"description,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputParam *AnimatedGraphicsOutputParam `json:"output_param,omitempty"`
}

func (o AnimatedGraphicsTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnimatedGraphicsTask struct{}"
	}

	return strings.Join([]string{"AnimatedGraphicsTask", string(data)}, " ")
}
