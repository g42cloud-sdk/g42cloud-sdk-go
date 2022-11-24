package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CacheConfig struct {
	IgnoreUrlParameter *bool `json:"ignore_url_parameter,omitempty"`

	FollowOrigin *bool `json:"follow_origin,omitempty"`

	Compress *CompressResponse `json:"compress,omitempty"`

	Rules *[]Rules `json:"rules,omitempty"`
}

func (o CacheConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheConfig struct{}"
	}

	return strings.Join([]string{"CacheConfig", string(data)}, " ")
}
