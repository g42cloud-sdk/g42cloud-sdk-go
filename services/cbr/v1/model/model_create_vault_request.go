package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVaultRequest struct {
	Body *VaultCreateReq `json:"body,omitempty"`
}

func (o CreateVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultRequest struct{}"
	}

	return strings.Join([]string{"CreateVaultRequest", string(data)}, " ")
}
