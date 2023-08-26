package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRemuxTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskResponse", string(data)}, " ")
}
