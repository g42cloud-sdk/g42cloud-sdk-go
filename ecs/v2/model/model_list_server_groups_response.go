package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListServerGroupsResponse struct {
	ServerGroups *[]ListServerGroupsResult `json:"server_groups,omitempty"`

	PageInfo       *ListServerGroupsPageInfoResult `json:"page_info,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListServerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListServerGroupsResponse", string(data)}, " ")
}
