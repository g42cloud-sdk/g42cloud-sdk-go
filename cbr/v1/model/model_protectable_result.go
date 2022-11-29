package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableResult struct {
	Code *string `json:"code,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Result bool `json:"result"`

	Vault *VaultGet `json:"vault,omitempty"`

	Message *string `json:"message,omitempty"`
}

func (o ProtectableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableResult struct{}"
	}

	return strings.Join([]string{"ProtectableResult", string(data)}, " ")
}
