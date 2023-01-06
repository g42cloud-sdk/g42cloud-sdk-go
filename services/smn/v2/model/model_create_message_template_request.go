package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMessageTemplateRequest struct {
	Body *CreateMessageTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateMessageTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateMessageTemplateRequest", string(data)}, " ")
}
