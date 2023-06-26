package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTemplateGroupRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *TransTemplateGroup `json:"body,omitempty"`
}

func (o CreateTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupRequest", string(data)}, " ")
}
