package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateExtractTaskRequest struct {
	Body *CreateExtractTaskReq `json:"body,omitempty"`
}

func (o CreateExtractTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtractTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateExtractTaskRequest", string(data)}, " ")
}
