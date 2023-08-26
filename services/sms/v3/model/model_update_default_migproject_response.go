package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDefaultMigprojectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDefaultMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDefaultMigprojectResponse", string(data)}, " ")
}
