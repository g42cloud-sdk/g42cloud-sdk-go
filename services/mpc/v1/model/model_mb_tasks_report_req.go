package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MbTasksReportReq struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	Retry *bool `json:"retry,omitempty"`

	Parameter *MbTaskParameter `json:"parameter,omitempty"`
}

func (o MbTasksReportReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MbTasksReportReq struct{}"
	}

	return strings.Join([]string{"MbTasksReportReq", string(data)}, " ")
}
