package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TaskByServerSource struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`

	State *string `json:"state,omitempty"`

	StartDate *int64 `json:"start_date,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	VmTemplateId *string `json:"vm_template_id,omitempty"`

	RegionId *string `json:"region_id,omitempty"`

	ProjectName *string `json:"project_name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	TargetServer *TargetServerById `json:"target_server,omitempty"`

	LogCollectStatus *string `json:"log_collect_status,omitempty"`

	ExistServer *bool `json:"exist_server,omitempty"`

	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`
}

func (o TaskByServerSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskByServerSource struct{}"
	}

	return strings.Join([]string{"TaskByServerSource", string(data)}, " ")
}
