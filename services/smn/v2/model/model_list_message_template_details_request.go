package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMessageTemplateDetailsRequest struct {
	MessageTemplateId string `json:"message_template_id"`
}

func (o ListMessageTemplateDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplateDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListMessageTemplateDetailsRequest", string(data)}, " ")
}
