package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotifySmnTopicConfigRequest struct {
	Body *NotificationConfigReq `json:"body,omitempty"`
}

func (o NotifySmnTopicConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifySmnTopicConfigRequest struct{}"
	}

	return strings.Join([]string{"NotifySmnTopicConfigRequest", string(data)}, " ")
}
