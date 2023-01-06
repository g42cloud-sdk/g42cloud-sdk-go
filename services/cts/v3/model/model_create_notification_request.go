package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateNotificationRequest struct {
	Body *CreateNotificationRequestBody `json:"body,omitempty"`
}

func (o CreateNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationRequest struct{}"
	}

	return strings.Join([]string{"CreateNotificationRequest", string(data)}, " ")
}
