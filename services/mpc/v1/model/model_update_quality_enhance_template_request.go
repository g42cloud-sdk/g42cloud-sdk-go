package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateQualityEnhanceTemplateRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *UpdateQualityEnhanceTemplateReq `json:"body,omitempty"`
}

func (o UpdateQualityEnhanceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateRequest", string(data)}, " ")
}
