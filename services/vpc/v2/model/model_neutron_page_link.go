package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronPageLink struct {
	Href string `json:"href"`

	Rel string `json:"rel"`
}

func (o NeutronPageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronPageLink struct{}"
	}

	return strings.Join([]string{"NeutronPageLink", string(data)}, " ")
}
