package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	TopicCount *int32 `json:"topic_count,omitempty"`

	Topics         *[]ListTopicsItem `json:"topics,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListTopicsResponse", string(data)}, " ")
}
