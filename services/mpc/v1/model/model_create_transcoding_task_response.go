package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTranscodingTaskResponse struct {
	TaskId         *int32 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateTranscodingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodingTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTranscodingTaskResponse", string(data)}, " ")
}
