package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Nics struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Cidr string `json:"cidr"`

	Ip *string `json:"ip,omitempty"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
