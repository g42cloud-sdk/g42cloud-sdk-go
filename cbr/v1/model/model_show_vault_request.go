package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVaultRequest struct {
	VaultId string `json:"vault_id"`
}

func (o ShowVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultRequest", string(data)}, " ")
}
