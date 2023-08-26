package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcTagsResponse struct {
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcTagsResponse", string(data)}, " ")
}
