package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteServerNicsResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteServerNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicsResponse", string(data)}, " ")
}
