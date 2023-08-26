package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteWatermarkTemplateRequest struct {
	TemplateId int32 `json:"template_id"`
}

func (o DeleteWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkTemplateRequest", string(data)}, " ")
}
