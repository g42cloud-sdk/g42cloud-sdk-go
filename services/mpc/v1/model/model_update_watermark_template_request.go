package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateWatermarkTemplateRequest struct {
	Body *WatermarkTemplate `json:"body,omitempty"`
}

func (o UpdateWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkTemplateRequest", string(data)}, " ")
}
