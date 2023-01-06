package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateAndDeleteVaultTagsRequest struct {
	VaultId string `json:"vault_id"`

	Body *BulkCreateAndDeleteVaultTagsReq `json:"body,omitempty"`
}

func (o BatchCreateAndDeleteVaultTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAndDeleteVaultTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsRequest", string(data)}, " ")
}
