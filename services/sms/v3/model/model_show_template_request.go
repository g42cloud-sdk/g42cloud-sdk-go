package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTemplateRequest struct {
	Id string `json:"id"`
}

func (o ShowTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateRequest", string(data)}, " ")
}
