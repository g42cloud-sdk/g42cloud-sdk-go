package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVaultResponse struct {
	Vaults *[]Vault `json:"vaults,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultResponse struct{}"
	}

	return strings.Join([]string{"ListVaultResponse", string(data)}, " ")
}
