package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
