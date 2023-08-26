package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTopicRequestBody struct {
	Name string `json:"name"`

	DisplayName string `json:"display_name"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateTopicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTopicRequestBody", string(data)}, " ")
}
