package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceGroupInfo struct {
	GroupName *string `json:"group_name,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	InstanceStatistics *InstanceStatistics `json:"instance_statistics,omitempty"`

	Status *string `json:"status,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ResourceGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupInfo struct{}"
	}

	return strings.Join([]string{"ResourceGroupInfo", string(data)}, " ")
}
