package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPortsResponse struct {
	Ports          *[]Port `json:"ports,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsResponse struct{}"
	}

	return strings.Join([]string{"ListPortsResponse", string(data)}, " ")
}
