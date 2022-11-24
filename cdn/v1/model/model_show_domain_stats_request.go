package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDomainStatsRequest struct {
	Action string `json:"action"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	DomainName string `json:"domain_name"`

	StatType string `json:"stat_type"`

	Interval *int64 `json:"interval,omitempty"`

	GroupBy *string `json:"group_by,omitempty"`

	ServiceArea *string `json:"service_area,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsRequest", string(data)}, " ")
}
