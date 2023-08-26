package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EngineExternalEntrypoint struct {
	ExternalAddress *string `json:"external_address,omitempty"`

	PublicAddress *string `json:"public_address,omitempty"`

	ServiceEndpoint map[string]EntrypointItem `json:"service_endpoint,omitempty"`

	PublicServiceEndpoint map[string]EntrypointItem `json:"public_service_endpoint,omitempty"`
}

func (o EngineExternalEntrypoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineExternalEntrypoint struct{}"
	}

	return strings.Join([]string{"EngineExternalEntrypoint", string(data)}, " ")
}
