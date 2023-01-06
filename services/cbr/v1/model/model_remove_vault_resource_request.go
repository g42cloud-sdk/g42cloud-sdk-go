package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RemoveVaultResourceRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultRemoveResourceReq `json:"body,omitempty"`
}

func (o RemoveVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"RemoveVaultResourceRequest", string(data)}, " ")
}
