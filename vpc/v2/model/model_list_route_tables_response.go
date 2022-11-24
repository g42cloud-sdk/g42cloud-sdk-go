package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRouteTablesResponse struct {
	Routetables    *[]RouteTableListResp `json:"routetables,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRouteTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRouteTablesResponse struct{}"
	}

	return strings.Join([]string{"ListRouteTablesResponse", string(data)}, " ")
}
