package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
