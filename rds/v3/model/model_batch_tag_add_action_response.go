package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchTagAddActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagAddActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagAddActionResponse struct{}"
	}

	return strings.Join([]string{"BatchTagAddActionResponse", string(data)}, " ")
}
