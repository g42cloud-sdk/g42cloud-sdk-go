package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddVaultResourceResponse struct {
	AddResourceIds *[]string `json:"add_resource_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"AddVaultResourceResponse", string(data)}, " ")
}
