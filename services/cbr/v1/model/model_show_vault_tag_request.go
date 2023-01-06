package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVaultTagRequest struct {
	VaultId string `json:"vault_id"`
}

func (o ShowVaultTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultTagRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultTagRequest", string(data)}, " ")
}
