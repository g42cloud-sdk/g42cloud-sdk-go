package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTaskStatusRequest struct {
	TaskId string `json:"task_id"`

	Body *UpdateTaskStatusReq `json:"body,omitempty"`
}

func (o UpdateTaskStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusRequest", string(data)}, " ")
}
