package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEncryptTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEncryptTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateEncryptTaskResponse", string(data)}, " ")
}
