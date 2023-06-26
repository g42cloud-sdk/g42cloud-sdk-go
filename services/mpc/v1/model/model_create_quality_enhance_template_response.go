package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateQualityEnhanceTemplateResponse struct {
	TemplateId     *int32 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateQualityEnhanceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateQualityEnhanceTemplateResponse", string(data)}, " ")
}
