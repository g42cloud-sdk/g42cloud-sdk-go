package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteVaultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultResponse struct{}"
	}

	return strings.Join([]string{"DeleteVaultResponse", string(data)}, " ")
}
