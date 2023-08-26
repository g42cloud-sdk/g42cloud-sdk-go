package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CancelRemuxTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o CancelRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelRemuxTaskRequest", string(data)}, " ")
}
