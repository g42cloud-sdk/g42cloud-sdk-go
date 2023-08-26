package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteQualityEnhanceTemplateRequest struct {
	TemplateId int32 `json:"template_id"`
}

func (o DeleteQualityEnhanceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteQualityEnhanceTemplateRequest", string(data)}, " ")
}
