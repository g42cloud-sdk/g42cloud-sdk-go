package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateResponse", string(data)}, " ")
}
