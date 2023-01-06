package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMessageTemplateRequest struct {
	MessageTemplateId string `json:"message_template_id"`

	Body *UpdateMessageTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateMessageTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateMessageTemplateRequest", string(data)}, " ")
}
