package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTopUrlRequest struct {
	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	DomainName string `json:"domain_name"`

	StatType string `json:"stat_type"`

	ServiceArea *string `json:"service_area,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowTopUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopUrlRequest struct{}"
	}

	return strings.Join([]string{"ShowTopUrlRequest", string(data)}, " ")
}
