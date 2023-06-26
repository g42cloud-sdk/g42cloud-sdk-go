package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEditingJobResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEditingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobResponse struct{}"
	}

	return strings.Join([]string{"CreateEditingJobResponse", string(data)}, " ")
}
