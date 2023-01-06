package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowResourceGroupResponse struct {
	GroupName *string `json:"group_name,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	Resources *[]ResourceGroup `json:"resources,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	MetaData *MetaData `json:"meta_data,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupResponse", string(data)}, " ")
}
