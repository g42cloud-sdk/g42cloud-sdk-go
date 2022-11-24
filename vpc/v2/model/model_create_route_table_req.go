package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRouteTableReq struct {
	Name *string `json:"name,omitempty"`

	Routes *[]RouteTableRoute `json:"routes,omitempty"`

	VpcId string `json:"vpc_id"`

	Description *string `json:"description,omitempty"`
}

func (o CreateRouteTableReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTableReq struct{}"
	}

	return strings.Join([]string{"CreateRouteTableReq", string(data)}, " ")
}
