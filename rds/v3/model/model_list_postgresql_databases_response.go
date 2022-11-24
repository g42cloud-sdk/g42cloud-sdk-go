package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPostgresqlDatabasesResponse struct {
	Databases *[]PostgresqlListDatabase `json:"databases,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPostgresqlDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabasesResponse", string(data)}, " ")
}
