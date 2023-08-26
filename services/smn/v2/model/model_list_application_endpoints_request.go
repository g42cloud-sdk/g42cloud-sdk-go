package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationEndpointsRequest struct {
	ApplicationUrn string `json:"application_urn"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Enabled *string `json:"enabled,omitempty"`

	Token *string `json:"token,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

func (o ListApplicationEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointsRequest", string(data)}, " ")
}
