package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultCreateReq struct {
	Vault *VaultCreate `json:"vault"`
}

func (o VaultCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreateReq struct{}"
	}

	return strings.Join([]string{"VaultCreateReq", string(data)}, " ")
}
