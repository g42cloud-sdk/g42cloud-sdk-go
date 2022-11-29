package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListBackupsResponse struct {
	Backups *[]BackupResp `json:"backups,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit          *int32 `json:"limit,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
