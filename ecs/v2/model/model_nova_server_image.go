package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaServerImage struct {
	Id string `json:"id"`

	Links []NovaLink `json:"links"`
}

func (o NovaServerImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerImage struct{}"
	}

	return strings.Join([]string{"NovaServerImage", string(data)}, " ")
}
