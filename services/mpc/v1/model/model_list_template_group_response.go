package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTemplateGroupResponse struct {
	TemplateGroupList *[]TemplateGroup `json:"template_group_list,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupResponse", string(data)}, " ")
}
