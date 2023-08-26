package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MessageTemplate struct {
	MessageTemplateId string `json:"message_template_id"`

	MessageTemplateName string `json:"message_template_name"`

	Protocol string `json:"protocol"`

	TagNames []string `json:"tag_names"`

	CreateTime string `json:"create_time"`

	UpdateTime string `json:"update_time"`
}

func (o MessageTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessageTemplate struct{}"
	}

	return strings.Join([]string{"MessageTemplate", string(data)}, " ")
}
