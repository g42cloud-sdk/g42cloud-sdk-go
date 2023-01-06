package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotificationBeginTime struct {
}

func (o NotificationBeginTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationBeginTime struct{}"
	}

	return strings.Join([]string{"NotificationBeginTime", string(data)}, " ")
}
