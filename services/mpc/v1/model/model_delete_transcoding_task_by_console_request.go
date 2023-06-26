package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTranscodingTaskByConsoleRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId int32 `json:"task_id"`
}

func (o DeleteTranscodingTaskByConsoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskByConsoleRequest struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskByConsoleRequest", string(data)}, " ")
}
