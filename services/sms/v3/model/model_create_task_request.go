package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTaskRequest struct {
	Body *PostTask `json:"body,omitempty"`
}

func (o CreateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskRequest", string(data)}, " ")
}
