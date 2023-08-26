package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVaultProjectTagRequest struct {
}

func (o ShowVaultProjectTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultProjectTagRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultProjectTagRequest", string(data)}, " ")
}
