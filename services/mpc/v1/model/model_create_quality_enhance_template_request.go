package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateQualityEnhanceTemplateRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *QualityEnhanceTemplate `json:"body,omitempty"`
}

func (o CreateQualityEnhanceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateQualityEnhanceTemplateRequest", string(data)}, " ")
}
