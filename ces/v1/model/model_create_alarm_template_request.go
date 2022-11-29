package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAlarmTemplateRequest struct {
	ContentType string `json:"Content-Type"`

	Body *CreateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateRequest", string(data)}, " ")
}
