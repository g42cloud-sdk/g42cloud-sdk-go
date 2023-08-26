package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmTemplatesResponse struct {
	AlarmTemplates *[]AlarmTemplate `json:"alarm_templates,omitempty"`

	MetaData       *MetaData `json:"meta_data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAlarmTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesResponse", string(data)}, " ")
}
