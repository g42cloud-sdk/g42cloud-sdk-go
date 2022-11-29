package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddVaultResourceRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultAddResourceReq `json:"body,omitempty"`
}

func (o AddVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"AddVaultResourceRequest", string(data)}, " ")
}
