package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMessageTemplateResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	MessageTemplateId *string `json:"message_template_id,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o CreateMessageTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateMessageTemplateResponse", string(data)}, " ")
}
