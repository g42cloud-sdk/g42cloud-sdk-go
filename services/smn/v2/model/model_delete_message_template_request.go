package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMessageTemplateRequest struct {
	MessageTemplateId string `json:"message_template_id"`
}

func (o DeleteMessageTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteMessageTemplateRequest", string(data)}, " ")
}
