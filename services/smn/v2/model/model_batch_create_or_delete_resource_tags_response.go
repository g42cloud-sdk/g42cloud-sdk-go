package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteResourceTagsResponse", string(data)}, " ")
}
