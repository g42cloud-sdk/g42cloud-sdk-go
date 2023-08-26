package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTemplateRequest struct {
	Body *CreateTemplateReq `json:"body,omitempty"`
}

func (o CreateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequest", string(data)}, " ")
}
