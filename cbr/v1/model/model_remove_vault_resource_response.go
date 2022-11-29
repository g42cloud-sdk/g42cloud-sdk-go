package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveVaultResourceResponse struct {
	RemoveResourceIds *[]string `json:"remove_resource_ids,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o RemoveVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"RemoveVaultResourceResponse", string(data)}, " ")
}
