package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicDetailsResponse struct {
	UpdateTime *string `json:"update_time,omitempty"`

	PushPolicy *int32 `json:"push_policy,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	Name *string `json:"name,omitempty"`

	TopicUrn *string `json:"topic_urn,omitempty"`

	DisplayName *string `json:"display_name,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListTopicDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListTopicDetailsResponse", string(data)}, " ")
}
