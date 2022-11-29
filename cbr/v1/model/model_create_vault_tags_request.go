package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
