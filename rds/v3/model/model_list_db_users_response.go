package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDbUsersResponse struct {
	Users *[]UserForList `json:"users,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDbUsersResponse", string(data)}, " ")
}
