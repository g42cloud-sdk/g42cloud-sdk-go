package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
