package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
