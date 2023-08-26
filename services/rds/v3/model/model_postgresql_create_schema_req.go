package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlCreateSchemaReq struct {
	SchemaName string `json:"schema_name"`

	Owner string `json:"owner"`
}

func (o PostgresqlCreateSchemaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlCreateSchemaReq struct{}"
	}

	return strings.Join([]string{"PostgresqlCreateSchemaReq", string(data)}, " ")
}
