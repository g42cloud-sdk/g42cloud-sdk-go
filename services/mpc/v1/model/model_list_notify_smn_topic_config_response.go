package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListNotifySmnTopicConfigResponse struct {
	Notifications *[]Notification `json:"notifications,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNotifySmnTopicConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotifySmnTopicConfigResponse struct{}"
	}

	return strings.Join([]string{"ListNotifySmnTopicConfigResponse", string(data)}, " ")
}
