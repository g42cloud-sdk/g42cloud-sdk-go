package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationEndpointAttributesRequest struct {
	EndpointUrn string `json:"endpoint_urn"`
}

func (o ListApplicationEndpointAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesRequest", string(data)}, " ")
}
