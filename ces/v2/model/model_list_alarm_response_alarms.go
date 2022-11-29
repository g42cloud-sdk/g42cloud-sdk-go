package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmResponseAlarms struct {
	AlarmId *string `json:"alarm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	Policies *[]Policy `json:"policies,omitempty"`

	Resources *[]ResourcesInListResp `json:"resources,omitempty"`

	Type *string `json:"type,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`
}

func (o ListAlarmResponseAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmResponseAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmResponseAlarms", string(data)}, " ")
}
