package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreheatingTaskRequestBody struct {
	Urls []string `json:"urls"`
}

func (o PreheatingTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreheatingTaskRequestBody struct{}"
	}

	return strings.Join([]string{"PreheatingTaskRequestBody", string(data)}, " ")
}
