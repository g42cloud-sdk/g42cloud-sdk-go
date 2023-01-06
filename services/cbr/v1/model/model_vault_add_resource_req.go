package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultAddResourceReq struct {
	Resources []ResourceCreate `json:"resources"`
}

func (o VaultAddResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultAddResourceReq struct{}"
	}

	return strings.Join([]string{"VaultAddResourceReq", string(data)}, " ")
}
