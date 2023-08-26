package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListNotificationsResponse struct {
	Notifications  *[]NotificationsResponseBody `json:"notifications,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationsResponse", string(data)}, " ")
}
