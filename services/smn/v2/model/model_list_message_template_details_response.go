package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMessageTemplateDetailsResponse struct {
	MessageTemplateId *string `json:"message_template_id,omitempty"`

	MessageTemplateName *string `json:"message_template_name,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	TagNames *[]string `json:"tag_names,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime *string `json:"update_time,omitempty"`

	Content *string `json:"content,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMessageTemplateDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplateDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTemplateDetailsResponse", string(data)}, " ")
}
