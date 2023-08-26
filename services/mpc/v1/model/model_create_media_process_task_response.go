package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMediaProcessTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMediaProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessTaskResponse", string(data)}, " ")
}
