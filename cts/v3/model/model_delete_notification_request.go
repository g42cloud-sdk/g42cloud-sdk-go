package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteNotificationRequest struct {
	NotificationId string `json:"notification_id"`
}

func (o DeleteNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotificationRequest", string(data)}, " ")
}
