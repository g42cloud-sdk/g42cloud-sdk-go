package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Quotas struct {
	QuotaLimit *int32 `json:"quota_limit,omitempty"`

	Type *string `json:"type,omitempty"`

	Used *int32 `json:"used,omitempty"`

	UserDomainId *string `json:"user_domain_id,omitempty"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
