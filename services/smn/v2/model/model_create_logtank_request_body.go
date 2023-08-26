package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateLogtankRequestBody struct {
	LogGroupId string `json:"log_group_id"`

	LogStreamId string `json:"log_stream_id"`
}

func (o CreateLogtankRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogtankRequestBody", string(data)}, " ")
}
