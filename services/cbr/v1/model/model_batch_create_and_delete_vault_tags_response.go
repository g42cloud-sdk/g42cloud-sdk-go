package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateAndDeleteVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateAndDeleteVaultTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAndDeleteVaultTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsResponse", string(data)}, " ")
}
