package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMessageTemplatesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	MessageTemplateCount *int32 `json:"message_template_count,omitempty"`

	MessageTemplates *[]MessageTemplate `json:"message_templates,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListMessageTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTemplatesResponse", string(data)}, " ")
}
