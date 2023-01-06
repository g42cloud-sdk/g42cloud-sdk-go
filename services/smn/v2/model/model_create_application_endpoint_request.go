package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateApplicationEndpointRequest struct {
	ApplicationUrn string `json:"application_urn"`

	Body *CreateApplicationEndpointRequestBody `json:"body,omitempty"`
}

func (o CreateApplicationEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationEndpointRequest", string(data)}, " ")
}
