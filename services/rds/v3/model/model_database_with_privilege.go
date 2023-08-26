package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DatabaseWithPrivilege struct {
	Name string `json:"name"`

	Readonly bool `json:"readonly"`
}

func (o DatabaseWithPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseWithPrivilege struct{}"
	}

	return strings.Join([]string{"DatabaseWithPrivilege", string(data)}, " ")
}
