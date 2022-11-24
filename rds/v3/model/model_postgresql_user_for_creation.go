package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostgresqlUserForCreation struct {
	Name string `json:"name"`

	Password string `json:"password"`
}

func (o PostgresqlUserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlUserForCreation struct{}"
	}

	return strings.Join([]string{"PostgresqlUserForCreation", string(data)}, " ")
}
