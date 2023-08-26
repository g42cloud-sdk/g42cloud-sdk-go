package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SgObject struct {
	Id string `json:"id"`

	Name string `json:"name"`
}

func (o SgObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SgObject struct{}"
	}

	return strings.Join([]string{"SgObject", string(data)}, " ")
}
