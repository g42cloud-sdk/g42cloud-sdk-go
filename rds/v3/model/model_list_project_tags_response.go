package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProjectTagsResponse struct {
	Tags           *[]ProjectTagInfoResponse `json:"tags,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTagsResponse", string(data)}, " ")
}
