package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTranscodingTaskByConsoleRequest struct {
	TaskId int32 `json:"task_id"`
}

func (o DeleteTranscodingTaskByConsoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskByConsoleRequest struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskByConsoleRequest", string(data)}, " ")
}
