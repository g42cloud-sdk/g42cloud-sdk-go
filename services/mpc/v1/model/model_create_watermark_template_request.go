package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateWatermarkTemplateRequest struct {
	Body *WatermarkTemplate `json:"body,omitempty"`
}

func (o CreateWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateWatermarkTemplateRequest", string(data)}, " ")
}
