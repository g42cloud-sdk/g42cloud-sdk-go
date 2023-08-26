package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteWatermarkTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkTemplateResponse", string(data)}, " ")
}
