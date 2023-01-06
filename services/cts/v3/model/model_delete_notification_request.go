package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
