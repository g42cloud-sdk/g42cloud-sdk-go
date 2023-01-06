package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TaskExecutorBrief struct {
	Duration *int64 `json:"duration,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o TaskExecutorBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskExecutorBrief struct{}"
	}

	return strings.Join([]string{"TaskExecutorBrief", string(data)}, " ")
}
