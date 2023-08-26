package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicAttributesResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTopicAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributesResponse", string(data)}, " ")
}
