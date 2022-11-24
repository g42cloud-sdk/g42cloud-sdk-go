package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecyclePolicy struct {
	RetentionPeriodInDays *string `json:"retention_period_in_days,omitempty"`
}

func (o RecyclePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicy struct{}"
	}

	return strings.Join([]string{"RecyclePolicy", string(data)}, " ")
}
