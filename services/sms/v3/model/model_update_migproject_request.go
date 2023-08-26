package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMigprojectRequest struct {
	MigProjectId string `json:"mig_project_id"`

	Body *MigProject `json:"body,omitempty"`
}

func (o UpdateMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMigprojectRequest struct{}"
	}

	return strings.Join([]string{"UpdateMigprojectRequest", string(data)}, " ")
}
