package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicAttributeByNameResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTopicAttributeByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributeByNameResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributeByNameResponse", string(data)}, " ")
}
