package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainItemDetailsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	DomainName string `json:"domain_name"`

	ServiceArea *string `json:"service_area,omitempty"`

	StatType string `json:"stat_type"`
}

func (o ShowDomainItemDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainItemDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainItemDetailsRequest", string(data)}, " ")
}
