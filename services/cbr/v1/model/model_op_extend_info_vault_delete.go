package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OpExtendInfoVaultDelete struct {
	FailCount *int32 `json:"fail_count,omitempty"`

	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o OpExtendInfoVaultDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoVaultDelete struct{}"
	}

	return strings.Join([]string{"OpExtendInfoVaultDelete", string(data)}, " ")
}
