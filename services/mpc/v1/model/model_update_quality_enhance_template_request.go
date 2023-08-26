package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateQualityEnhanceTemplateRequest struct {
	Body *UpdateQualityEnhanceTemplateReq `json:"body,omitempty"`
}

func (o UpdateQualityEnhanceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateRequest", string(data)}, " ")
}
