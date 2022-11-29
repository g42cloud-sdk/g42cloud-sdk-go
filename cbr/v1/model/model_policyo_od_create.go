package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyoOdCreate struct {
	DayBackups *int32 `json:"day_backups,omitempty"`

	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	DestinationRegion *string `json:"destination_region,omitempty"`

	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`

	MaxBackups *int32 `json:"max_backups,omitempty"`

	MonthBackups *int32 `json:"month_backups,omitempty"`

	RetentionDurationDays *int32 `json:"retention_duration_days,omitempty"`

	Timezone *string `json:"timezone,omitempty"`

	WeekBackups *int32 `json:"week_backups,omitempty"`

	YearBackups *int32 `json:"year_backups,omitempty"`
}

func (o PolicyoOdCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyoOdCreate struct{}"
	}

	return strings.Join([]string{"PolicyoOdCreate", string(data)}, " ")
}
