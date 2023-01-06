package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicRequestBody struct {
	DisplayName string `json:"display_name"`
}

func (o UpdateTopicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTopicRequestBody", string(data)}, " ")
}
