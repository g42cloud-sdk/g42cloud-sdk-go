package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupResponse struct {
	Backup         *BackupResp `json:"backup,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupResponse", string(data)}, " ")
}
