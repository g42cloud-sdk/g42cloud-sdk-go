package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchAddOrDeleteTagsRequest struct {
	ImageId string `json:"image_id"`

	Body *BatchAddOrDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchAddOrDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsRequest", string(data)}, " ")
}
