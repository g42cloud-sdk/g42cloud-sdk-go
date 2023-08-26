package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTemplateReq struct {
	Template *TemplateRequest `json:"template"`
}

func (o CreateTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateTemplateReq", string(data)}, " ")
}
