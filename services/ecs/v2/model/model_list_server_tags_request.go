package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListServerTagsRequest struct {
}

func (o ListServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListServerTagsRequest", string(data)}, " ")
}