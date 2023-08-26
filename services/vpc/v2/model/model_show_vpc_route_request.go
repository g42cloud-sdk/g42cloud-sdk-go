package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcRouteRequest struct {
	RouteId string `json:"route_id"`
}

func (o ShowVpcRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcRouteRequest", string(data)}, " ")
}
