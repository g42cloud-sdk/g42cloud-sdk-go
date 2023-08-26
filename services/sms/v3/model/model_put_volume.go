package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PutVolume struct {
	Id *string `json:"id,omitempty"`

	NeedMigration *bool `json:"need_migration,omitempty"`

	AdjustSize *int64 `json:"adjust_size,omitempty"`
}

func (o PutVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutVolume struct{}"
	}

	return strings.Join([]string{"PutVolume", string(data)}, " ")
}
