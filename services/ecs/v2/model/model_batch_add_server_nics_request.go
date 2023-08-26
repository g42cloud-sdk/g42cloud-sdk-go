package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchAddServerNicsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchAddServerNicsRequestBody `json:"body,omitempty"`
}

func (o BatchAddServerNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerNicsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddServerNicsRequest", string(data)}, " ")
}
