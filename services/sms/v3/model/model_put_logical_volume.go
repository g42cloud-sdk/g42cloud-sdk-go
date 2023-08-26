package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PutLogicalVolume struct {
	Id string `json:"id"`

	NeedMigration *bool `json:"need_migration,omitempty"`

	AdjustSize *int64 `json:"adjust_size,omitempty"`
}

func (o PutLogicalVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutLogicalVolume struct{}"
	}

	return strings.Join([]string{"PutLogicalVolume", string(data)}, " ")
}
