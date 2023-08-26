package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRouteTableResponse struct{}"
	}

	return strings.Join([]string{"ShowRouteTableResponse", string(data)}, " ")
}
