package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaListServersDetailsResponse struct {
	Servers *[]NovaServer `json:"servers,omitempty"`

	ServersLinks   *[]PageLink `json:"servers_links,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o NovaListServersDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListServersDetailsResponse struct{}"
	}

	return strings.Join([]string{"NovaListServersDetailsResponse", string(data)}, " ")
}
