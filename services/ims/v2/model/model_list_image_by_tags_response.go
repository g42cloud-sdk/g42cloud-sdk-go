package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListImageByTagsResponse struct {
	Resources *[]ShowImageByTagsResource `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListImageByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImageByTagsResponse", string(data)}, " ")
}
