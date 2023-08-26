package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateResponseHeaderResponse struct {
	Headers        *HeaderMap `json:"headers,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o UpdateResponseHeaderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResponseHeaderResponse struct{}"
	}

	return strings.Join([]string{"UpdateResponseHeaderResponse", string(data)}, " ")
}
