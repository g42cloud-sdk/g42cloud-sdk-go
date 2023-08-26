package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrateServerOption struct {
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`
}

func (o MigrateServerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateServerOption struct{}"
	}

	return strings.Join([]string{"MigrateServerOption", string(data)}, " ")
}
