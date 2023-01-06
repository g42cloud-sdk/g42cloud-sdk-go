package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotificationUsers struct {
	UserGroup string `json:"user_group"`

	UserList []string `json:"user_list"`
}

func (o NotificationUsers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationUsers struct{}"
	}

	return strings.Join([]string{"NotificationUsers", string(data)}, " ")
}
