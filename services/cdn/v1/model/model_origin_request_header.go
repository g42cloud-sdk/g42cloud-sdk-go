package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OriginRequestHeader struct {
	Name string `json:"name"`

	Value *string `json:"value,omitempty"`

	Action string `json:"action"`
}

func (o OriginRequestHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginRequestHeader struct{}"
	}

	return strings.Join([]string{"OriginRequestHeader", string(data)}, " ")
}
