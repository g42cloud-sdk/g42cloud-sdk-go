package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTranscodingTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId int32 `json:"task_id"`
}

func (o DeleteTranscodingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskRequest", string(data)}, " ")
}
