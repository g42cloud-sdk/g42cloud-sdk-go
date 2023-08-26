package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
