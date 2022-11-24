package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostgresqlGrantRequest struct {
	DbName string `json:"db_name"`

	Users []PostgresqlUserWithPrivilege `json:"users"`
}

func (o PostgresqlGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlGrantRequest struct{}"
	}

	return strings.Join([]string{"PostgresqlGrantRequest", string(data)}, " ")
}
