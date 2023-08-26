package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceRequestBody struct {
	Tags *[]ResourceTags `json:"tags,omitempty"`

	TagsAny *[]ResourceTags `json:"tags_any,omitempty"`

	NotTags *[]ResourceTags `json:"not_tags,omitempty"`

	NotTagsAny *[]ResourceTags `json:"not_tags_any,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Action string `json:"action"`

	Matches *[]TagMatch `json:"matches,omitempty"`
}

func (o ListInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstanceRequestBody", string(data)}, " ")
}
