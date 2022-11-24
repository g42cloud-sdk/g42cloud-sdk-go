package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreExistInstanceResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreExistInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreExistInstanceResponse", string(data)}, " ")
}
