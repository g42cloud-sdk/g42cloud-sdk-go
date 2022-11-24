package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Route struct {
	Destination *string `json:"destination,omitempty"`

	Nexthop *string `json:"nexthop,omitempty"`
}

func (o Route) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Route struct{}"
	}

	return strings.Join([]string{"Route", string(data)}, " ")
}
