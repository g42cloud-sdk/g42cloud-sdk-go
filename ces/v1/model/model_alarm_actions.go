package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmActions struct {
	Type string `json:"type"`

	NotificationList []string `json:"notificationList"`
}

func (o AlarmActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmActions struct{}"
	}

	return strings.Join([]string{"AlarmActions", string(data)}, " ")
}
