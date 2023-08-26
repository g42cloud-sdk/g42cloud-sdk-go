package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultResponse", string(data)}, " ")
}
