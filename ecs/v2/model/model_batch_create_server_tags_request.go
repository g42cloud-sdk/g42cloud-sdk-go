package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateServerTagsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchCreateServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateServerTagsRequest", string(data)}, " ")
}
