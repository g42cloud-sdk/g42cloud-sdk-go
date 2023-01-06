package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteAlarmTemplateRequest struct {
	ContentType string `json:"Content-Type"`

	TemplateId string `json:"template_id"`
}

func (o DeleteAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateRequest", string(data)}, " ")
}
