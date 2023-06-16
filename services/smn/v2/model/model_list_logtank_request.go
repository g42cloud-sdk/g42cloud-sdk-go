package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListLogtankRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o ListLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtankRequest struct{}"
	}

	return strings.Join([]string{"ListLogtankRequest", string(data)}, " ")
}
