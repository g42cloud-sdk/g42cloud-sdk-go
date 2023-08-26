package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RegisterServerRequest struct {
	Body *PostSourceServerBody `json:"body,omitempty"`
}

func (o RegisterServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerRequest struct{}"
	}

	return strings.Join([]string{"RegisterServerRequest", string(data)}, " ")
}
