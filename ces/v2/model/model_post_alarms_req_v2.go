package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostAlarmsReqV2 struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Namespace string `json:"namespace"`

	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	Resources [][]Dimension `json:"resources"`

	Policies []Policy `json:"policies"`

	Type string `json:"type"`

	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Enabled bool `json:"enabled"`

	NotificationEnabled bool `json:"notification_enabled"`

	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`
}

func (o PostAlarmsReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAlarmsReqV2 struct{}"
	}

	return strings.Join([]string{"PostAlarmsReqV2", string(data)}, " ")
}
