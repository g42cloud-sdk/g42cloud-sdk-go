package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDomainItemLocationDetailsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	DomainName string `json:"domain_name"`

	StatType string `json:"stat_type"`

	Region string `json:"region"`

	Isp string `json:"isp"`
}

func (o ShowDomainItemLocationDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainItemLocationDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainItemLocationDetailsRequest", string(data)}, " ")
}
