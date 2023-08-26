package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTemplatesResponse struct {
	Count *int32 `json:"count,omitempty"`

	Templates      *[]TemplateResponseBody `json:"templates,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
