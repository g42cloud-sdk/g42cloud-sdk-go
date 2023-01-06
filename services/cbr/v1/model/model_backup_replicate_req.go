package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BackupReplicateReq struct {
	Replicate *BackupReplicateReqBody `json:"replicate"`
}

func (o BackupReplicateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupReplicateReq struct{}"
	}

	return strings.Join([]string{"BackupReplicateReq", string(data)}, " ")
}
