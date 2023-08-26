package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type SourceServersResponseBody struct {
	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Name *string `json:"name,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	AddDate *int64 `json:"add_date,omitempty"`

	OsType *SourceServersResponseBodyOsType `json:"os_type,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`

	OemSystem *bool `json:"oem_system,omitempty"`

	State *SourceServersResponseBodyState `json:"state,omitempty"`

	Connected *bool `json:"connected,omitempty"`

	CpuQuantity *int32 `json:"cpu_quantity,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	CurrentTask *TaskByServerSources `json:"current_task,omitempty"`

	Checks *[]EnvironmentCheck `json:"checks,omitempty"`

	InitTargetServer *InitTargetServer `json:"init_target_server,omitempty"`

	Replicatesize *int64 `json:"replicatesize,omitempty"`

	StageActionTime *int64 `json:"stage_action_time,omitempty"`

	Totalsize *int64 `json:"totalsize,omitempty"`

	LastVisitTime *int64 `json:"last_visit_time,omitempty"`

	MigrationCycle *SourceServersResponseBodyMigrationCycle `json:"migration_cycle,omitempty"`

	StateActionTime *int64 `json:"state_action_time,omitempty"`
}

func (o SourceServersResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceServersResponseBody struct{}"
	}

	return strings.Join([]string{"SourceServersResponseBody", string(data)}, " ")
}

type SourceServersResponseBodyOsType struct {
	value string
}

type SourceServersResponseBodyOsTypeEnum struct {
	WINDOWS SourceServersResponseBodyOsType
	LINUX   SourceServersResponseBodyOsType
}

func GetSourceServersResponseBodyOsTypeEnum() SourceServersResponseBodyOsTypeEnum {
	return SourceServersResponseBodyOsTypeEnum{
		WINDOWS: SourceServersResponseBodyOsType{
			value: "WINDOWS",
		},
		LINUX: SourceServersResponseBodyOsType{
			value: "LINUX",
		},
	}
}

func (c SourceServersResponseBodyOsType) Value() string {
	return c.value
}

func (c SourceServersResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServersResponseBodyOsType) UnmarshalJSON(b []byte) error {
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

type SourceServersResponseBodyState struct {
	value string
}

type SourceServersResponseBodyStateEnum struct {
	UNAVAILABLE SourceServersResponseBodyState
	WAITING     SourceServersResponseBodyState
	INITIALIZE  SourceServersResponseBodyState
	REPLICATE   SourceServersResponseBodyState
	SYNCING     SourceServersResponseBodyState
	STOPPING    SourceServersResponseBodyState
	STOPPED     SourceServersResponseBodyState
	DELETING    SourceServersResponseBodyState
	ERROR       SourceServersResponseBodyState
	CLONING     SourceServersResponseBodyState
	CUTOVERING  SourceServersResponseBodyState
	FINISHED    SourceServersResponseBodyState
}

func GetSourceServersResponseBodyStateEnum() SourceServersResponseBodyStateEnum {
	return SourceServersResponseBodyStateEnum{
		UNAVAILABLE: SourceServersResponseBodyState{
			value: "unavailable",
		},
		WAITING: SourceServersResponseBodyState{
			value: "waiting",
		},
		INITIALIZE: SourceServersResponseBodyState{
			value: "initialize",
		},
		REPLICATE: SourceServersResponseBodyState{
			value: "replicate",
		},
		SYNCING: SourceServersResponseBodyState{
			value: "syncing",
		},
		STOPPING: SourceServersResponseBodyState{
			value: "stopping",
		},
		STOPPED: SourceServersResponseBodyState{
			value: "stopped",
		},
		DELETING: SourceServersResponseBodyState{
			value: "deleting",
		},
		ERROR: SourceServersResponseBodyState{
			value: "error",
		},
		CLONING: SourceServersResponseBodyState{
			value: "cloning",
		},
		CUTOVERING: SourceServersResponseBodyState{
			value: "cutovering",
		},
		FINISHED: SourceServersResponseBodyState{
			value: "finished",
		},
	}
}

func (c SourceServersResponseBodyState) Value() string {
	return c.value
}

func (c SourceServersResponseBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServersResponseBodyState) UnmarshalJSON(b []byte) error {
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

type SourceServersResponseBodyMigrationCycle struct {
	value string
}

type SourceServersResponseBodyMigrationCycleEnum struct {
	CUTOVERING  SourceServersResponseBodyMigrationCycle
	CUTOVERED   SourceServersResponseBodyMigrationCycle
	CHECKING    SourceServersResponseBodyMigrationCycle
	SETTING     SourceServersResponseBodyMigrationCycle
	REPLICATING SourceServersResponseBodyMigrationCycle
	SYNCING     SourceServersResponseBodyMigrationCycle
}

func GetSourceServersResponseBodyMigrationCycleEnum() SourceServersResponseBodyMigrationCycleEnum {
	return SourceServersResponseBodyMigrationCycleEnum{
		CUTOVERING: SourceServersResponseBodyMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: SourceServersResponseBodyMigrationCycle{
			value: "cutovered",
		},
		CHECKING: SourceServersResponseBodyMigrationCycle{
			value: "checking",
		},
		SETTING: SourceServersResponseBodyMigrationCycle{
			value: "setting",
		},
		REPLICATING: SourceServersResponseBodyMigrationCycle{
			value: "replicating",
		},
		SYNCING: SourceServersResponseBodyMigrationCycle{
			value: "syncing",
		},
	}
}

func (c SourceServersResponseBodyMigrationCycle) Value() string {
	return c.value
}

func (c SourceServersResponseBodyMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServersResponseBodyMigrationCycle) UnmarshalJSON(b []byte) error {
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
