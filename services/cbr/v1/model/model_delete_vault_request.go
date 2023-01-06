package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteVaultRequest struct {
	VaultId string `json:"vault_id"`
}

func (o DeleteVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultRequest struct{}"
	}

	return strings.Join([]string{"DeleteVaultRequest", string(data)}, " ")
}
