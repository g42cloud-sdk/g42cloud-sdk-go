package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListImageTagsRequest struct {
	ImageId string `json:"image_id"`
}

func (o ListImageTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImageTagsRequest", string(data)}, " ")
}
