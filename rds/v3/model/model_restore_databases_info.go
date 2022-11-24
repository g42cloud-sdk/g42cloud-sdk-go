package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreDatabasesInfo struct {
	Database string `json:"database"`

	Tables []RestoreTableInfo `json:"tables"`
}

func (o RestoreDatabasesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDatabasesInfo struct{}"
	}

	return strings.Join([]string{"RestoreDatabasesInfo", string(data)}, " ")
}
