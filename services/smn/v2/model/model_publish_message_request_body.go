package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublishMessageRequestBody struct {
	Subject *string `json:"subject,omitempty"`

	Message *string `json:"message,omitempty"`

	MessageStructure *string `json:"message_structure,omitempty"`

	MessageTemplateName *string `json:"message_template_name,omitempty"`

	Tags map[string]string `json:"tags,omitempty"`

	TimeToLive *string `json:"time_to_live,omitempty"`
}

func (o PublishMessageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishMessageRequestBody struct{}"
	}

	return strings.Join([]string{"PublishMessageRequestBody", string(data)}, " ")
}
