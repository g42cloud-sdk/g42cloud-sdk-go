package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPostgresqlDbUserPaginatedResponse struct {
	Users *[]PostgresqlUserForList `json:"users,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPostgresqlDbUserPaginatedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresqlDbUserPaginatedResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDbUserPaginatedResponse", string(data)}, " ")
}
