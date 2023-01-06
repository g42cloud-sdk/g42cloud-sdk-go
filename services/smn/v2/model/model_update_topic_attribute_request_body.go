package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicAttributeRequestBody struct {
	Value string `json:"value"`
}

func (o UpdateTopicAttributeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAttributeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTopicAttributeRequestBody", string(data)}, " ")
}
