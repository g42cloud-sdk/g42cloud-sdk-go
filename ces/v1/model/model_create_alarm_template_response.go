package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAlarmTemplateResponse struct {
	TemplateId     *string `json:"template_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateResponse", string(data)}, " ")
}
