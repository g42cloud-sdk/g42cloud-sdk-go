package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteServerNicsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchDeleteServerNicsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteServerNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicsRequest", string(data)}, " ")
}
