package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVpcTagsResponse struct {
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcTagsResponse", string(data)}, " ")
}
