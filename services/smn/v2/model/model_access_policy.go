package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AccessPolicy struct {
	Version string `json:"Version"`

	Id string `json:"Id"`

	Statement []Statement `json:"Statement"`
}

func (o AccessPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicy struct{}"
	}

	return strings.Join([]string{"AccessPolicy", string(data)}, " ")
}
