package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListImagesTagsRequest struct {
}

func (o ListImagesTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImagesTagsRequest", string(data)}, " ")
}
