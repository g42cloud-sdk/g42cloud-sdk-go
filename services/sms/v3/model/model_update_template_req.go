package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTemplateReq struct {
	Template *TemplateRequest `json:"template,omitempty"`
}

func (o UpdateTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateReq struct{}"
	}

	return strings.Join([]string{"UpdateTemplateReq", string(data)}, " ")
}
