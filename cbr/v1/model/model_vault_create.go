package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultCreate struct {
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	Billing *BillingCreate `json:"billing"`

	Description *string `json:"description,omitempty"`

	Name string `json:"name"`

	Resources []ResourceCreate `json:"resources"`

	Tags *[]Tag `json:"tags,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	AutoExpand *bool `json:"auto_expand,omitempty"`
}

func (o VaultCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreate struct{}"
	}

	return strings.Join([]string{"VaultCreate", string(data)}, " ")
}
