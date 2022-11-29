package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vault struct {
	Billing *Billing `json:"billing"`

	Description *string `json:"description,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	ProviderId string `json:"provider_id"`

	Resources []ResourceResp `json:"resources"`

	Tags *[]Tag `json:"tags,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	AutoExpand *bool `json:"auto_expand,omitempty"`

	SmnNotify *bool `json:"smn_notify,omitempty"`

	Threshold *int32 `json:"threshold,omitempty"`
}

func (o Vault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vault struct{}"
	}

	return strings.Join([]string{"Vault", string(data)}, " ")
}
