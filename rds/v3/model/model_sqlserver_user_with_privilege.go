package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlserverUserWithPrivilege struct {
	Name string `json:"name"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (o SqlserverUserWithPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverUserWithPrivilege struct{}"
	}

	return strings.Join([]string{"SqlserverUserWithPrivilege", string(data)}, " ")
}
