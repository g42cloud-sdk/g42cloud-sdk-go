package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotificationConfigReq struct {
	Notifications []Notification `json:"notifications"`
}

func (o NotificationConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationConfigReq struct{}"
	}

	return strings.Join([]string{"NotificationConfigReq", string(data)}, " ")
}
