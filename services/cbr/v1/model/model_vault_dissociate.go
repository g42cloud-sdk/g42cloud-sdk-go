package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultDissociate struct {
	PolicyId string `json:"policy_id"`
}

func (o VaultDissociate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultDissociate struct{}"
	}

	return strings.Join([]string{"VaultDissociate", string(data)}, " ")
}
