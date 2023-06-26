package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTemplateGroupResponse struct {
	TemplateGroup  *TemplateGroup `json:"template_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupResponse", string(data)}, " ")
}
