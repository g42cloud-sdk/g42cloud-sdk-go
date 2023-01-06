package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrateVaultResourceResponse struct {
	MigratedResources *[]string `json:"migrated_resources,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o MigrateVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceResponse", string(data)}, " ")
}
