package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVaultResponse struct{}"
	}

	return strings.Join([]string{"UpdateVaultResponse", string(data)}, " ")
}
