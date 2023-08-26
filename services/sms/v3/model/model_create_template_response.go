package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTemplateResponse struct {
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateResponse", string(data)}, " ")
}
