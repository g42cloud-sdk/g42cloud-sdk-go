package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAuthorizedDbUsersResponse struct {
	Users *[]UserWithPrivilege `json:"users,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuthorizedDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizedDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizedDbUsersResponse", string(data)}, " ")
}
