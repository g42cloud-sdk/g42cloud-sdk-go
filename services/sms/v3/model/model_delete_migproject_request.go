package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMigprojectRequest struct {
	MigProjectId string `json:"mig_project_id"`
}

func (o DeleteMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigprojectRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigprojectRequest", string(data)}, " ")
}
