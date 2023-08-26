package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplatesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplatesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesResponse", string(data)}, " ")
}
