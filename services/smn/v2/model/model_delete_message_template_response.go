package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMessageTemplateResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMessageTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteMessageTemplateResponse", string(data)}, " ")
}
