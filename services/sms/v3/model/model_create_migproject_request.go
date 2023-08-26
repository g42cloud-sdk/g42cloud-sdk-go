package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMigprojectRequest struct {
	Body *PostMigProjectBody `json:"body,omitempty"`
}

func (o CreateMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigprojectRequest struct{}"
	}

	return strings.Join([]string{"CreateMigprojectRequest", string(data)}, " ")
}
