package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VpcObject struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Cidr *string `json:"cidr,omitempty"`
}

func (o VpcObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcObject struct{}"
	}

	return strings.Join([]string{"VpcObject", string(data)}, " ")
}
