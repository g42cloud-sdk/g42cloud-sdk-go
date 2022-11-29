package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoRestore struct {
	BackupId *string `json:"backup_id,omitempty"`

	BackupName *string `json:"backup_name,omitempty"`

	TargetResourceId *string `json:"target_resource_id,omitempty"`

	TargetResourceName *string `json:"target_resource_name,omitempty"`
}

func (o OpExtendInfoRestore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoRestore struct{}"
	}

	return strings.Join([]string{"OpExtendInfoRestore", string(data)}, " ")
}
