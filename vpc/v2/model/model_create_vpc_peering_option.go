package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpcPeeringOption struct {
	Name string `json:"name"`

	RequestVpcInfo *VpcInfo `json:"request_vpc_info"`

	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info"`
}

func (o CreateVpcPeeringOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcPeeringOption struct{}"
	}

	return strings.Join([]string{"CreateVpcPeeringOption", string(data)}, " ")
}
