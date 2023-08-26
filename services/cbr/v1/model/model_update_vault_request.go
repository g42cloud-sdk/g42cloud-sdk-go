package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateVaultRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultUpdateReq `json:"body,omitempty"`
}

func (o UpdateVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVaultRequest struct{}"
	}

	return strings.Join([]string{"UpdateVaultRequest", string(data)}, " ")
}
