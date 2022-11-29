package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLoadBalancerRequestBody struct {
	Loadbalancer *CreateLoadBalancerOption `json:"loadbalancer"`
}

func (o CreateLoadBalancerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerRequestBody", string(data)}, " ")
}
