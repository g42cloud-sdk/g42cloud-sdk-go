package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEditingJobRequest struct {
	Body *CreateEditingJobReq `json:"body,omitempty"`
}

func (o CreateEditingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateEditingJobRequest", string(data)}, " ")
}
