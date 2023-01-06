package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteVaultTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVaultTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteVaultTagResponse", string(data)}, " ")
}
