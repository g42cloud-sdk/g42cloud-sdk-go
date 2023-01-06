package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrateFollowerResponse struct {
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateFollowerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateFollowerResponse struct{}"
	}

	return strings.Join([]string{"MigrateFollowerResponse", string(data)}, " ")
}
