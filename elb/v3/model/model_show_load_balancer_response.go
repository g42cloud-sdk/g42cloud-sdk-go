package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowLoadBalancerResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Loadbalancer   *LoadBalancer `json:"loadbalancer,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerResponse", string(data)}, " ")
}
