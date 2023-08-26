package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BasicTaskInfo struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o BasicTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicTaskInfo struct{}"
	}

	return strings.Join([]string{"BasicTaskInfo", string(data)}, " ")
}
