package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTargetPasswordResponse struct {
	TemplateId *string `json:"template_id,omitempty"`

	TargetPassword *string `json:"target_password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTargetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTargetPasswordResponse struct{}"
	}

	return strings.Join([]string{"ShowTargetPasswordResponse", string(data)}, " ")
}
