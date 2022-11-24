package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDomainLocationStatsRequest struct {
	Action string `json:"action"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	DomainName string `json:"domain_name"`

	StatType string `json:"stat_type"`

	Interval *int64 `json:"interval,omitempty"`

	Country *string `json:"country,omitempty"`

	Province *string `json:"province,omitempty"`

	Isp *string `json:"isp,omitempty"`

	GroupBy *string `json:"group_by,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainLocationStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainLocationStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainLocationStatsRequest", string(data)}, " ")
}
