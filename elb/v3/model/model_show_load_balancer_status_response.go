package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowLoadBalancerStatusResponse struct {
	Statuses *LoadBalancerStatusResult `json:"statuses,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLoadBalancerStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerStatusResponse", string(data)}, " ")
}
