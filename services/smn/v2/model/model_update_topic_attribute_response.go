package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicAttributeResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTopicAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAttributeResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicAttributeResponse", string(data)}, " ")
}
