package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateServerTagsResponse", string(data)}, " ")
}
