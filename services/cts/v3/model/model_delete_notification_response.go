package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotificationResponse", string(data)}, " ")
}
