package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTemplateRequest struct {
	Id string `json:"id"`

	Body *UpdateTemplateReq `json:"body,omitempty"`
}

func (o UpdateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateRequest", string(data)}, " ")
}
