package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVaultTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultTagsResponse", string(data)}, " ")
}
