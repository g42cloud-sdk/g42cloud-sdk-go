package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubnetsByTagsResponse struct {
	Resources *[]ListResourceResp `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsResponse", string(data)}, " ")
}
