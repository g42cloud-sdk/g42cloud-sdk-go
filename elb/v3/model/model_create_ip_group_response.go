package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateIpGroupResponse", string(data)}, " ")
}
