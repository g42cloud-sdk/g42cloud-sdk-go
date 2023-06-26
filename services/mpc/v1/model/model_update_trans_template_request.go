package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTransTemplateRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *ModifyTransTemplateReq `json:"body,omitempty"`
}

func (o UpdateTransTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTransTemplateRequest", string(data)}, " ")
}
