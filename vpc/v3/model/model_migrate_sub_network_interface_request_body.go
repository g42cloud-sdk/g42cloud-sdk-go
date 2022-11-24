package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateSubNetworkInterfaceRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	MigrationInfo *MigrateSubNetworkInterfaceOption `json:"migration_info"`
}

func (o MigrateSubNetworkInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateSubNetworkInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateSubNetworkInterfaceRequestBody", string(data)}, " ")
}
