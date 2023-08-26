package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateNetworkCheckInfoResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNetworkCheckInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkCheckInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateNetworkCheckInfoResponse", string(data)}, " ")
}
