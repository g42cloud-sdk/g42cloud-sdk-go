package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSubnetTagsResponse struct {
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSubnetTagsResponse", string(data)}, " ")
}
