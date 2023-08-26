package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateHttpsInfoResponse struct {
	Https          *HttpInfoResponseBody `json:"https,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateHttpsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpsInfoResponse", string(data)}, " ")
}
