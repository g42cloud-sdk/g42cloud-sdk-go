package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerResponse", string(data)}, " ")
}
