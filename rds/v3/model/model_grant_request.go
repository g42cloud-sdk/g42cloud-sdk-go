package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GrantRequest struct {
	DbName string `json:"db_name"`

	Users []UserWithPrivilege `json:"users"`
}

func (o GrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantRequest struct{}"
	}

	return strings.Join([]string{"GrantRequest", string(data)}, " ")
}
