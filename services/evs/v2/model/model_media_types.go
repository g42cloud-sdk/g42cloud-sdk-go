package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MediaTypes struct {
	Base string `json:"base"`

	Type string `json:"type"`
}

func (o MediaTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaTypes struct{}"
	}

	return strings.Join([]string{"MediaTypes", string(data)}, " ")
}
