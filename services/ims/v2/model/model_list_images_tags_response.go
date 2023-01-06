package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListImagesTagsResponse struct {
	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListImagesTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImagesTagsResponse", string(data)}, " ")
}
