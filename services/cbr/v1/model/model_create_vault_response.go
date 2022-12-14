package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVaultResponse struct {
	Vault          *VaultCreateResource `json:"vault,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultResponse", string(data)}, " ")
}
