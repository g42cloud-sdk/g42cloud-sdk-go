package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDefaultMigprojectRequest struct {
	MigProjectId string `json:"mig_project_id"`
}

func (o UpdateDefaultMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultMigprojectRequest struct{}"
	}

	return strings.Join([]string{"UpdateDefaultMigprojectRequest", string(data)}, " ")
}
