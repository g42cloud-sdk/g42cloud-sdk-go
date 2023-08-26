package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTemplateGroupRequest struct {
	Body *TransTemplateGroup `json:"body,omitempty"`
}

func (o CreateTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupRequest", string(data)}, " ")
}
