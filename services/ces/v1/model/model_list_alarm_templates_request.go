package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmTemplatesRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmTemplateId *string `json:"alarmTemplateId,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	Dname *string `json:"dname,omitempty"`

	Start *string `json:"start,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesRequest", string(data)}, " ")
}
