package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateResponse", string(data)}, " ")
}
