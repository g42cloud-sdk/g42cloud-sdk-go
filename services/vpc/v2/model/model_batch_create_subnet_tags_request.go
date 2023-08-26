package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateSubnetTagsRequest struct {
	SubnetId string `json:"subnet_id"`

	Body *BatchCreateSubnetTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSubnetTagsRequest", string(data)}, " ")
}
