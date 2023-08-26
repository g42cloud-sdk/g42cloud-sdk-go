package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplateRequest struct {
	TemplateId int64 `json:"template_id"`
}

func (o DeleteTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateRequest", string(data)}, " ")
}
