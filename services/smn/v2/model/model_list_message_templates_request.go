package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMessageTemplatesRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	MessageTemplateName *string `json:"message_template_name,omitempty"`

	Protocol *string `json:"protocol,omitempty"`
}

func (o ListMessageTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListMessageTemplatesRequest", string(data)}, " ")
}
