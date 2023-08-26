package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadBalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerRequest", string(data)}, " ")
}
