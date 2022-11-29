package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
