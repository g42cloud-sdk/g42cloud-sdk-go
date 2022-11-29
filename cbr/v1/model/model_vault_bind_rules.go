package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultBindRules struct {
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o VaultBindRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBindRules struct{}"
	}

	return strings.Join([]string{"VaultBindRules", string(data)}, " ")
}