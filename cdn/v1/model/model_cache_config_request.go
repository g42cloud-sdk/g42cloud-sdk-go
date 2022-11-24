package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CacheConfigRequest struct {
	IgnoreUrlParameter *bool `json:"ignore_url_parameter,omitempty"`

	FollowOrigin *bool `json:"follow_origin,omitempty"`

	Compress *CompressRequest `json:"compress,omitempty"`

	Rules *[]Rules `json:"rules,omitempty"`
}

func (o CacheConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheConfigRequest struct{}"
	}

	return strings.Join([]string{"CacheConfigRequest", string(data)}, " ")
}
