package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePriorityRequestBody struct {
	Id string `json:"id"`

	Priority int32 `json:"priority"`
}

func (o BatchUpdatePriorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePriorityRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePriorityRequestBody", string(data)}, " ")
}
