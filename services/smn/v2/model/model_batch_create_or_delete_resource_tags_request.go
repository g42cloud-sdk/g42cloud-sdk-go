package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteResourceTagsRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Body *BatchCreateOrDeleteResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteResourceTagsRequest", string(data)}, " ")
}
