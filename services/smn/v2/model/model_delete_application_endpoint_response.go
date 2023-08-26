package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteApplicationEndpointResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteApplicationEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationEndpointResponse", string(data)}, " ")
}
