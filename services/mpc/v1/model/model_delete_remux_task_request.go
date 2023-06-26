package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteRemuxTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId string `json:"task_id"`
}

func (o DeleteRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteRemuxTaskRequest", string(data)}, " ")
}
