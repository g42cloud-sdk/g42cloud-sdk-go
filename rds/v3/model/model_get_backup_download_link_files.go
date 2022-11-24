package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetBackupDownloadLinkFiles struct {
	Name string `json:"name"`

	Size int64 `json:"size"`

	DownloadLink string `json:"download_link"`

	LinkExpiredTime string `json:"link_expired_time"`

	DatabaseName string `json:"database_name"`
}

func (o GetBackupDownloadLinkFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetBackupDownloadLinkFiles struct{}"
	}

	return strings.Join([]string{"GetBackupDownloadLinkFiles", string(data)}, " ")
}
