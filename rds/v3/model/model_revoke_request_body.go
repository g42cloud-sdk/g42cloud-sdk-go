package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RevokeRequestBody struct {
	DbName string `json:"db_name"`

	Users []RevokeRequestBodyUsers `json:"users"`
}

func (o RevokeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeRequestBody struct{}"
	}

	return strings.Join([]string{"RevokeRequestBody", string(data)}, " ")
}
