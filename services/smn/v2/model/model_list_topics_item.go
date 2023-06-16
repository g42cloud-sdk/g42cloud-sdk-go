package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicsItem struct {
	TopicUrn string `json:"topic_urn"`

	Name string `json:"name"`

	DisplayName string `json:"display_name"`

	PushPolicy int32 `json:"push_policy"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	TopicId string `json:"topic_id"`
}

func (o ListTopicsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsItem struct{}"
	}

	return strings.Join([]string{"ListTopicsItem", string(data)}, " ")
}
