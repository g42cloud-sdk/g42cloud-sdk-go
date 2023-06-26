package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTransTemplateRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	XVodProjectId *string `json:"x-vod-projectId,omitempty"`

	Body *TransTemplate `json:"body,omitempty"`
}

func (o CreateTransTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTransTemplateRequest", string(data)}, " ")
}
