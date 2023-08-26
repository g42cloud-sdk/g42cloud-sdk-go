package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RestoreToExistingInstanceResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreToExistingInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceResponse", string(data)}, " ")
}
