package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTaskSpeedRequest struct {
	TaskId string `json:"task_id"`

	Body *UpdateTaskSpeedReq `json:"body,omitempty"`
}

func (o UpdateTaskSpeedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSpeedRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskSpeedRequest", string(data)}, " ")
}
