package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
