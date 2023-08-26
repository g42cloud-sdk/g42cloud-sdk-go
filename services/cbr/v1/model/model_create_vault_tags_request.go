package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVaultTagsRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultTagsCreateReq `json:"body,omitempty"`
}

func (o CreateVaultTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateVaultTagsRequest", string(data)}, " ")
}
