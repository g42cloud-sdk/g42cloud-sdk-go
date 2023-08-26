package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateResponse", string(data)}, " ")
}
