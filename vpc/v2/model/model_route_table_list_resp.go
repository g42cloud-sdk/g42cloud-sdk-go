package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RouteTableListResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Default bool `json:"default"`

	Subnets []SubnetList `json:"subnets"`

	TenantId string `json:"tenant_id"`

	VpcId string `json:"vpc_id"`

	Description string `json:"description"`
}

func (o RouteTableListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteTableListResp struct{}"
	}

	return strings.Join([]string{"RouteTableListResp", string(data)}, " ")
}
