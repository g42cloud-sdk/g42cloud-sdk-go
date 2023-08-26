package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateApplicationEndpointRequest struct {
	EndpointUrn string `json:"endpoint_urn"`

	Body *UpdateApplicationEndpointRequestBody `json:"body,omitempty"`
}

func (o UpdateApplicationEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationEndpointRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationEndpointRequest", string(data)}, " ")
}
