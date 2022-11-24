package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteServerNicsRequestBody struct {
	Nics []BatchDeleteServerNicOption `json:"nics"`
}

func (o BatchDeleteServerNicsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicsRequestBody", string(data)}, " ")
}
