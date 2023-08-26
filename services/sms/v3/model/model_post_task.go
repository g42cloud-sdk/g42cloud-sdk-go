package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PostTask struct {
	Name string `json:"name"`

	Type PostTaskType `json:"type"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	OsType string `json:"os_type"`

	SourceServer *SourceServerByTask `json:"source_server"`

	TargetServer *TargetServerByTask `json:"target_server"`

	MigrationIp *string `json:"migration_ip,omitempty"`

	RegionName string `json:"region_name"`

	RegionId string `json:"region_id"`

	ProjectName string `json:"project_name"`

	ProjectId string `json:"project_id"`

	VmTemplateId *string `json:"vm_template_id,omitempty"`

	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	Syncing *bool `json:"syncing,omitempty"`

	ExistServer *bool `json:"exist_server,omitempty"`

	StartNetworkCheck *bool `json:"start_network_check,omitempty"`
}

func (o PostTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostTask struct{}"
	}

	return strings.Join([]string{"PostTask", string(data)}, " ")
}

type PostTaskType struct {
	value string
}

type PostTaskTypeEnum struct {
	MIGRATE_FILE  PostTaskType
	MIGRATE_BLOCK PostTaskType
}

func GetPostTaskTypeEnum() PostTaskTypeEnum {
	return PostTaskTypeEnum{
		MIGRATE_FILE: PostTaskType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: PostTaskType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c PostTaskType) Value() string {
	return c.value
}

func (c PostTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostTaskType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
