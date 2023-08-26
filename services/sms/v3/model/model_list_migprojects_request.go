package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMigprojectsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListMigprojectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigprojectsRequest struct{}"
	}

	return strings.Join([]string{"ListMigprojectsRequest", string(data)}, " ")
}
