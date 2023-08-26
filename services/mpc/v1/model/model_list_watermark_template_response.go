package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListWatermarkTemplateResponse struct {
	Total *int32 `json:"total,omitempty"`

	Templates      *[]WatermarkTemplate `json:"templates,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateResponse", string(data)}, " ")
}
