package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowMigprojectRequest struct {
	MigProjectId string `json:"mig_project_id"`
}

func (o ShowMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigprojectRequest struct{}"
	}

	return strings.Join([]string{"ShowMigprojectRequest", string(data)}, " ")
}
