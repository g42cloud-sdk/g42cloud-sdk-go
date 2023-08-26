package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationEndpointAttributesResponseBodyAttributes struct {
	Enabled string `json:"enabled"`

	Token string `json:"token"`

	UserData string `json:"user_data"`
}

func (o ListApplicationEndpointAttributesResponseBodyAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesResponseBodyAttributes struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesResponseBodyAttributes", string(data)}, " ")
}
