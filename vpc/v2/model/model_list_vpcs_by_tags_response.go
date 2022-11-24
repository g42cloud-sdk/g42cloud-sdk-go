package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVpcsByTagsResponse struct {
	Resources *[]ListResourceResp `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVpcsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcsByTagsResponse", string(data)}, " ")
}
