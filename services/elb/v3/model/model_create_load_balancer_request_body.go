package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
