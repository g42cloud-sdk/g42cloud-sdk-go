package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrationErrors struct {
	ErrorJson *string `json:"error_json,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	Name *string `json:"name,omitempty"`

	SourceId *string `json:"source_id,omitempty"`

	SourceIp *string `json:"source_ip,omitempty"`

	TargetIp *string `json:"target_ip,omitempty"`
}

func (o MigrationErrors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationErrors struct{}"
	}

	return strings.Join([]string{"MigrationErrors", string(data)}, " ")
}
