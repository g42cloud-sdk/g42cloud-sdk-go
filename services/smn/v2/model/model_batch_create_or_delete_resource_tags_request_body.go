package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteResourceTagsRequestBody struct {
	Tags []ResourceTag `json:"tags"`

	Action string `json:"action"`
}

func (o BatchCreateOrDeleteResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteResourceTagsRequestBody", string(data)}, " ")
}
