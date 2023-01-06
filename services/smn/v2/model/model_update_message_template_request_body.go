package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMessageTemplateRequestBody struct {
	Content string `json:"content"`
}

func (o UpdateMessageTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMessageTemplateRequestBody", string(data)}, " ")
}
