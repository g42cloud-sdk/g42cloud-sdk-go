package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RegisterServerResponse struct {
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerResponse struct{}"
	}

	return strings.Join([]string{"RegisterServerResponse", string(data)}, " ")
}
