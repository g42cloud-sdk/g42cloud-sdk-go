package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TaskByServerSources struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`

	State *string `json:"state,omitempty"`

	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	StartDate *int64 `json:"start_date,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	CompressRate *float64 `json:"compress_rate,omitempty"`

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

	RemainSeconds *int64 `json:"remain_seconds,omitempty"`

	LogBucket *string `json:"log_bucket,omitempty"`

	LogExpire *int64 `json:"log_expire,omitempty"`

	LogUploadTime *int64 `json:"log_upload_time,omitempty"`

	LogShareUrl *string `json:"log_share_url,omitempty"`
}

func (o TaskByServerSources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskByServerSources struct{}"
	}

	return strings.Join([]string{"TaskByServerSources", string(data)}, " ")
}
