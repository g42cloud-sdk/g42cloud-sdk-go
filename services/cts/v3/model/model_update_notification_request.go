package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateNotificationRequest struct {
	Body *UpdateNotificationRequestBody `json:"body,omitempty"`
}

func (o UpdateNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationRequest", string(data)}, " ")
}
