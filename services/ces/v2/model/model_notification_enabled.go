package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotificationEnabled struct {
}

func (o NotificationEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationEnabled struct{}"
	}

	return strings.Join([]string{"NotificationEnabled", string(data)}, " ")
}
