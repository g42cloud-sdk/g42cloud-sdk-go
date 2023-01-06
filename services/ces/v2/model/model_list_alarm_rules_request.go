package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmRulesRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId *string `json:"alarm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesRequest", string(data)}, " ")
}
