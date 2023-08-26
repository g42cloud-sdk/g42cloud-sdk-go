package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTranscodingTaskRequest struct {
	Body *CreateTranscodingReq `json:"body,omitempty"`
}

func (o CreateTranscodingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodingTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTranscodingTaskRequest", string(data)}, " ")
}
