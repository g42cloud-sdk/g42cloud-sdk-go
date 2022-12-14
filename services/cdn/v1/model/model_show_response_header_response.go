package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowResponseHeaderResponse struct {
	Headers        *HeaderMap `json:"headers,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowResponseHeaderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResponseHeaderResponse struct{}"
	}

	return strings.Join([]string{"ShowResponseHeaderResponse", string(data)}, " ")
}
