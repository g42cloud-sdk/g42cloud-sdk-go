package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcRouteResponse struct {
	Route          *VpcRoute `json:"route,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowVpcRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcRouteResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcRouteResponse", string(data)}, " ")
}
