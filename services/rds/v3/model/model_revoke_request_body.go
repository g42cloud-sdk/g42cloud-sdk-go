package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
