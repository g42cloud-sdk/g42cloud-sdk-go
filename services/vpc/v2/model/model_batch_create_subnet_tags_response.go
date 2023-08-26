package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSubnetTagsResponse", string(data)}, " ")
}
