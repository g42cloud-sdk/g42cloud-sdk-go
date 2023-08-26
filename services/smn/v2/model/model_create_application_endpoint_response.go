package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateApplicationEndpointResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	EndpointUrn    *string `json:"endpoint_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApplicationEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationEndpointResponse", string(data)}, " ")
}
