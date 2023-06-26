package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateQualityEnhanceTemplateReq struct {
	TemplateId *int32 `json:"template_id,omitempty"`

	Template *QualityEnhanceTemplate `json:"template,omitempty"`
}

func (o UpdateQualityEnhanceTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateReq struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateReq", string(data)}, " ")
}
