package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultUpdateReq struct {
	Vault *VaultUpdate `json:"vault"`
}

func (o VaultUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultUpdateReq struct{}"
	}

	return strings.Join([]string{"VaultUpdateReq", string(data)}, " ")
}
