package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicAttributesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Attributes     *TopicAttribute `json:"attributes,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTopicAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListTopicAttributesResponse", string(data)}, " ")
}
