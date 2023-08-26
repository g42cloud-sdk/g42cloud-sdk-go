package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMigprojectsResponse struct {
	Count *int32 `json:"count,omitempty"`

	Migprojects    *[]MigprojectsResponseBody `json:"migprojects,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListMigprojectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigprojectsResponse struct{}"
	}

	return strings.Join([]string{"ListMigprojectsResponse", string(data)}, " ")
}
