package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMigprojectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigprojectResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigprojectResponse", string(data)}, " ")
}
