package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmTemplateRequest struct {
	ContentType string `json:"Content-Type"`

	TemplateId string `json:"template_id"`

	Body *UpdateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequest", string(data)}, " ")
}
