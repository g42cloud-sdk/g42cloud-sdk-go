package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTemplateGroupRequest struct {
	Body *ModifyTransTemplateGroup `json:"body,omitempty"`
}

func (o UpdateTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupRequest", string(data)}, " ")
}
