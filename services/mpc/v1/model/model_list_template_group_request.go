package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTemplateGroupRequest struct {
	GroupId *[]string `json:"group_id,omitempty"`

	GroupName *[]string `json:"group_name,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupRequest", string(data)}, " ")
}
