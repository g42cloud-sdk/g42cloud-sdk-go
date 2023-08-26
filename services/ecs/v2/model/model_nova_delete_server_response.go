package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaDeleteServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDeleteServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteServerResponse struct{}"
	}

	return strings.Join([]string{"NovaDeleteServerResponse", string(data)}, " ")
}
