package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListServerTagsResponse struct {
	Tags           *[]ProjectTag `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerTagsResponse struct{}"
	}

	return strings.Join([]string{"ListServerTagsResponse", string(data)}, " ")
}
