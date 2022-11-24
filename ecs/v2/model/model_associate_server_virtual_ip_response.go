package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateServerVirtualIpResponse struct {
	PortId         *string `json:"port_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateServerVirtualIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpResponse struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpResponse", string(data)}, " ")
}
