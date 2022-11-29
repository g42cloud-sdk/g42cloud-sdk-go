package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultUpdate struct {
	Billing *BillingUpdate `json:"billing,omitempty"`

	Name *string `json:"name,omitempty"`

	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	AutoExpand *bool `json:"auto_expand,omitempty"`

	SmnNotify *bool `json:"smn_notify,omitempty"`

	Threshold *int32 `json:"threshold,omitempty"`
}

func (o VaultUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultUpdate struct{}"
	}

	return strings.Join([]string{"VaultUpdate", string(data)}, " ")
}
