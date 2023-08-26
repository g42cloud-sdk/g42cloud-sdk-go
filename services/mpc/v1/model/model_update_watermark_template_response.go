package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateWatermarkTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkTemplateResponse", string(data)}, " ")
}
