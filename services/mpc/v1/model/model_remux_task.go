package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RemuxTask struct {
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

	OutputParam *RemuxOutputParam `json:"output_param,omitempty"`

	CompleteRatio *int32 `json:"complete_ratio,omitempty"`

	OutputMetadata *MetaData `json:"output_metadata,omitempty"`
}

func (o RemuxTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemuxTask struct{}"
	}

	return strings.Join([]string{"RemuxTask", string(data)}, " ")
}
