package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSubnetTagsRequest struct {
}

func (o ListSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetTagsRequest", string(data)}, " ")
}
