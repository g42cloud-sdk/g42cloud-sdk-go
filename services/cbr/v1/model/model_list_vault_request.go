package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListVaultRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	CloudType *string `json:"cloud_type,omitempty"`

	ProtectType *string `json:"protect_type,omitempty"`

	ObjectType *string `json:"object_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Id *string `json:"id,omitempty"`

	PolicyId *string `json:"policy_id,omitempty"`

	Status *string `json:"status,omitempty"`

	ResourceIds *string `json:"resource_ids,omitempty"`
}

func (o ListVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultRequest struct{}"
	}

	return strings.Join([]string{"ListVaultRequest", string(data)}, " ")
}
