package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEncryptTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId string `json:"task_id"`
}

func (o DeleteEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskRequest", string(data)}, " ")
}
