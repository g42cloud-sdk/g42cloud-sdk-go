package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupPolicy struct {
	KeepDays int32 `json:"keep_days"`

	StartTime *string `json:"start_time,omitempty"`

	Period *string `json:"period,omitempty"`
}

func (o BackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicy struct{}"
	}

	return strings.Join([]string{"BackupPolicy", string(data)}, " ")
}
