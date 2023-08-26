package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListErrorServersResponse struct {
	Count *int32 `json:"count,omitempty"`

	MigrationErrors *[]MigrationErrors `json:"migration_errors,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ListErrorServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorServersResponse struct{}"
	}

	return strings.Join([]string{"ListErrorServersResponse", string(data)}, " ")
}
