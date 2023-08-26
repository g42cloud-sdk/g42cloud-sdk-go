package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTemplateResponse struct {
	Template       *TemplateResponseBody `json:"template,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateResponse", string(data)}, " ")
}
