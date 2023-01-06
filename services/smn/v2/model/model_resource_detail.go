package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResourceDetail struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`

	DetailId string `json:"detailId"`

	TopicUrn string `json:"topic_urn"`

	DisplayName string `json:"display_name"`
}

func (o ResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDetail struct{}"
	}

	return strings.Join([]string{"ResourceDetail", string(data)}, " ")
}
