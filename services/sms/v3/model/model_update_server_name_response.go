package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerNameResponse", string(data)}, " ")
}
