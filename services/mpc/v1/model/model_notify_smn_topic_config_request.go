package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotifySmnTopicConfigRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *NotificationConfigReq `json:"body,omitempty"`
}

func (o NotifySmnTopicConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifySmnTopicConfigRequest struct{}"
	}

	return strings.Join([]string{"NotifySmnTopicConfigRequest", string(data)}, " ")
}
