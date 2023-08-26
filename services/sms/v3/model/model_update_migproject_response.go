package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMigprojectResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateMigprojectResponse", string(data)}, " ")
}
