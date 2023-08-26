package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVaultResourceInstancesRequest struct {
	Body *VaultResourceInstancesReq `json:"body,omitempty"`
}

func (o ShowVaultResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultResourceInstancesRequest", string(data)}, " ")
}
