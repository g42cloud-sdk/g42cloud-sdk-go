package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationEndpointAttributesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Attributes     *ListApplicationEndpointAttributesResponseBodyAttributes `json:"attributes,omitempty"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o ListApplicationEndpointAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesResponse", string(data)}, " ")
}
