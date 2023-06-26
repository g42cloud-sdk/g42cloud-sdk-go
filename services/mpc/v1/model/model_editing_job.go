package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EditingJob struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	Description *string `json:"description,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	EditType *[]string `json:"edit_type,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	EditTaskReq *CreateEditingJobReq `json:"edit_task_req,omitempty"`

	OutputFileInfo *[]OutputFileInfo `json:"output_file_info,omitempty"`
}

func (o EditingJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingJob struct{}"
	}

	return strings.Join([]string{"EditingJob", string(data)}, " ")
}
