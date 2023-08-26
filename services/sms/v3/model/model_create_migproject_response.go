package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMigprojectResponse struct {
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigprojectResponse struct{}"
	}

	return strings.Join([]string{"CreateMigprojectResponse", string(data)}, " ")
}
