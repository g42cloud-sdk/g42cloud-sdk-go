package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteVaultTagRequest struct {
	Key string `json:"key"`

	VaultId string `json:"vault_id"`
}

func (o DeleteVaultTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteVaultTagRequest", string(data)}, " ")
}
