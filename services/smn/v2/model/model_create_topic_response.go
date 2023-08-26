package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTopicResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateTopicResponse", string(data)}, " ")
}
