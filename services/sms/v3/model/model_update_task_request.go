package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTaskRequest struct {
	TaskId string `json:"task_id"`

	Body *PutTaskReq `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
