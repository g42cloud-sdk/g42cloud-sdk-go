package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationEndpointsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	NextPageFlag *bool `json:"next_page_flag,omitempty"`

	Endpoints      *[]ApplicationEndpoint `json:"endpoints,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListApplicationEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointsResponse", string(data)}, " ")
}
