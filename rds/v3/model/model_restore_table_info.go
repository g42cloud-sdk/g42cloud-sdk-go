package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTableInfo struct {
	OldName string `json:"oldName"`

	NewName string `json:"newName"`
}

func (o RestoreTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTableInfo struct{}"
	}

	return strings.Join([]string{"RestoreTableInfo", string(data)}, " ")
}
