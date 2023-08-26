package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteTaskRequest", string(data)}, " ")
}
