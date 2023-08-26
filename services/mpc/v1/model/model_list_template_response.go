package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTemplateResponse struct {
	TemplateArray *[]TemplateInfo `json:"template_array,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateResponse", string(data)}, " ")
}
