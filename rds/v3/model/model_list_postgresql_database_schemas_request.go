package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPostgresqlDatabaseSchemasRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db_name"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`
}

func (o ListPostgresqlDatabaseSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabaseSchemasRequest struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabaseSchemasRequest", string(data)}, " ")
}
