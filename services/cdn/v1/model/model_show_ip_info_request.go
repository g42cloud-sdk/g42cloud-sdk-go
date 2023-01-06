package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
