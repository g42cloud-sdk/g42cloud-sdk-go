package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateResponse", string(data)}, " ")
}
