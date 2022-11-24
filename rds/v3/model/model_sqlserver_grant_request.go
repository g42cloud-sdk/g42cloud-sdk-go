package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
