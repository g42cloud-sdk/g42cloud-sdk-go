package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateLoadBalancerRequest struct {
	Body *CreateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerRequest", string(data)}, " ")
}
