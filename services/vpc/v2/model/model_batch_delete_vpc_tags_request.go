package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteVpcTagsRequest struct {
	VpcId string `json:"vpc_id"`

	Body *BatchDeleteVpcTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteVpcTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpcTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpcTagsRequest", string(data)}, " ")
}
