package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteResetTracksTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId string `json:"task_id"`
}

func (o DeleteResetTracksTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskRequest", string(data)}, " ")
}
