package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteQualityEnhanceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteQualityEnhanceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteQualityEnhanceTemplateResponse", string(data)}, " ")
}
