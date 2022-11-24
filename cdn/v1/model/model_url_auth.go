package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UrlAuth struct {
	Status string `json:"status"`

	Type *string `json:"type,omitempty"`

	Key *string `json:"key,omitempty"`

	TimeFormat *string `json:"time_format,omitempty"`

	ExpireTime *int32 `json:"expire_time,omitempty"`
}

func (o UrlAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlAuth struct{}"
	}

	return strings.Join([]string{"UrlAuth", string(data)}, " ")
}
