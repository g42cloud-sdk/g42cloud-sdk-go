package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowIpGroupResponse", string(data)}, " ")
}
