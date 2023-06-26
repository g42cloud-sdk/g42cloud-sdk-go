package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTransTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTransTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTransTemplateResponse", string(data)}, " ")
}
