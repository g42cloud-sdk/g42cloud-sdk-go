package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationAttributesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	ApplicationId *string `json:"application_id,omitempty"`

	Attributes     *ListApplicationAttributesResponseBodyAttributes `json:"attributes,omitempty"`
	HttpStatusCode int                                              `json:"-"`
}

func (o ListApplicationAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationAttributesResponse", string(data)}, " ")
}
