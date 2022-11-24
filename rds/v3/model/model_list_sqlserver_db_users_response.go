package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSqlserverDbUsersResponse struct {
	Users *[]UserForList `json:"users,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlserverDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlserverDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListSqlserverDbUsersResponse", string(data)}, " ")
}
