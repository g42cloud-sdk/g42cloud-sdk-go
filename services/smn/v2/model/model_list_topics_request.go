package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicsRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Name *string `json:"name,omitempty"`

	FuzzyName *string `json:"fuzzy_name,omitempty"`

	TopicId *string `json:"topic_id,omitempty"`
}

func (o ListTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListTopicsRequest", string(data)}, " ")
}
