package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaCreateKeypairRequest struct {
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`

	Body *NovaCreateKeypairRequestBody `json:"body,omitempty"`
}

func (o NovaCreateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairRequest", string(data)}, " ")
}
