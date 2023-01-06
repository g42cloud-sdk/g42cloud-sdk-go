package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ReplicationRecordsExtraInfo struct {
	Progress *int32 `json:"progress,omitempty"`

	FailCode *string `json:"fail_code,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`

	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	DestinatioVaultId *string `json:"destinatio_vault_id,omitempty"`
}

func (o ReplicationRecordsExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationRecordsExtraInfo struct{}"
	}

	return strings.Join([]string{"ReplicationRecordsExtraInfo", string(data)}, " ")
}
