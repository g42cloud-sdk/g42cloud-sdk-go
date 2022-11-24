package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchTagActionDelRequestBody struct {
	Action string `json:"action"`

	Tags []TagDelWithKeyValue `json:"tags"`
}

func (o BatchTagActionDelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionDelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionDelRequestBody", string(data)}, " ")
}
