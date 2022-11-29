package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoadBalancerRef struct {
	Id *string `json:"id,omitempty"`
}

func (o LoadBalancerRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerRef struct{}"
	}

	return strings.Join([]string{"LoadBalancerRef", string(data)}, " ")
}
