package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEncryptTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskRequest", string(data)}, " ")
}
