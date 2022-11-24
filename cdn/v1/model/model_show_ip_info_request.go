package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowIpInfoRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Ips string `json:"ips"`
}

func (o ShowIpInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowIpInfoRequest", string(data)}, " ")
}
