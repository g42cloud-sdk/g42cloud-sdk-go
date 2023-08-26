package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultBindRules struct {
	Tags *[]BindRulesTags `json:"tags,omitempty"`
}

func (o VaultBindRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBindRules struct{}"
	}

	return strings.Join([]string{"VaultBindRules", string(data)}, " ")
}
