package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateQualityEnhanceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateQualityEnhanceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateResponse", string(data)}, " ")
}
