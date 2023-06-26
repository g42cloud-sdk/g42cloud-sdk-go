package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CommonCreateTaskRsp struct {
	TaskId *string `json:"task_id,omitempty"`
}

func (o CommonCreateTaskRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonCreateTaskRsp struct{}"
	}

	return strings.Join([]string{"CommonCreateTaskRsp", string(data)}, " ")
}
