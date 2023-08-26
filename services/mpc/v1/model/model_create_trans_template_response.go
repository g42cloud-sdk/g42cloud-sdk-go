package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTransTemplateResponse struct {
	TemplateId     *int32 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateTransTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateTransTemplateResponse", string(data)}, " ")
}
