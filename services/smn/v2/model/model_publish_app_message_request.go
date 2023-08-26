package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublishAppMessageRequest struct {
	EndpointUrn string `json:"endpoint_urn"`

	Body *PublishAppMessageRequestBody `json:"body,omitempty"`
}

func (o PublishAppMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppMessageRequest struct{}"
	}

	return strings.Join([]string{"PublishAppMessageRequest", string(data)}, " ")
}
