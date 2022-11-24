package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePostgresqlDatabaseSchemaRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *PostgresqlDatabaseSchemaReq `json:"body,omitempty"`
}

func (o CreatePostgresqlDatabaseSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseSchemaRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseSchemaRequest", string(data)}, " ")
}
