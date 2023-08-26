package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SqlserverGrantRequest struct {
	DbName string `json:"db_name"`

	Users []SqlserverUserWithPrivilege `json:"users"`
}

func (o SqlserverGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverGrantRequest struct{}"
	}

	return strings.Join([]string{"SqlserverGrantRequest", string(data)}, " ")
}
