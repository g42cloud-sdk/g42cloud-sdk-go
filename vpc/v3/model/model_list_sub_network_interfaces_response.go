package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubNetworkInterfacesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterfaces *[]SubNetworkInterface `json:"sub_network_interfaces,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSubNetworkInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfacesResponse", string(data)}, " ")
}
