package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowHttpInfoResponse struct {
	Https          *HttpInfoResponseBody `json:"https,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowHttpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpInfoResponse", string(data)}, " ")
}
