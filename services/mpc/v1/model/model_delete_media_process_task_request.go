package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMediaProcessTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteMediaProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMediaProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMediaProcessTaskRequest", string(data)}, " ")
}
