package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowServerResponse struct {
	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Name *string `json:"name,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	AddDate *int64 `json:"add_date,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`

	OemSystem *bool `json:"oem_system,omitempty"`

	State *ShowServerResponseState `json:"state,omitempty"`

	Connected *bool `json:"connected,omitempty"`

	Firmware *ShowServerResponseFirmware `json:"firmware,omitempty"`

	InitTargetServer *InitTargetServer `json:"init_target_server,omitempty"`

	CpuQuantity *int32 `json:"cpu_quantity,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	CurrentTask *TaskByServerSource `json:"current_task,omitempty"`

	Disks *[]ServerDisk `json:"disks,omitempty"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`

	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty"`

	Networks *[]NetWork `json:"networks,omitempty"`

	Checks *[]EnvironmentCheck `json:"checks,omitempty"`

	MigrationCycle *ShowServerResponseMigrationCycle `json:"migration_cycle,omitempty"`

	StateActionTime *int64 `json:"state_action_time,omitempty"`

	Replicatesize *int64 `json:"replicatesize,omitempty"`

	Totalsize *int64 `json:"totalsize,omitempty"`

	LastVisitTime *int64 `json:"last_visit_time,omitempty"`

	StageActionTime *int64 `json:"stage_action_time,omitempty"`

	AgentVersion   *string `json:"agent_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerResponse struct{}"
	}

	return strings.Join([]string{"ShowServerResponse", string(data)}, " ")
}

type ShowServerResponseState struct {
	value string
}

type ShowServerResponseStateEnum struct {
	UNAVAILABLE ShowServerResponseState
	WAITING     ShowServerResponseState
	INITIALIZE  ShowServerResponseState
	REPLICATE   ShowServerResponseState
	SYNCING     ShowServerResponseState
	STOPPING    ShowServerResponseState
	STOPPED     ShowServerResponseState
	DELETING    ShowServerResponseState
	ERROR       ShowServerResponseState
	CLONING     ShowServerResponseState
	TESTING     ShowServerResponseState
	FINISHED    ShowServerResponseState
}

func GetShowServerResponseStateEnum() ShowServerResponseStateEnum {
	return ShowServerResponseStateEnum{
		UNAVAILABLE: ShowServerResponseState{
			value: "unavailable",
		},
		WAITING: ShowServerResponseState{
			value: "waiting",
		},
		INITIALIZE: ShowServerResponseState{
			value: "initialize",
		},
		REPLICATE: ShowServerResponseState{
			value: "replicate",
		},
		SYNCING: ShowServerResponseState{
			value: "syncing",
		},
		STOPPING: ShowServerResponseState{
			value: "stopping",
		},
		STOPPED: ShowServerResponseState{
			value: "stopped",
		},
		DELETING: ShowServerResponseState{
			value: "deleting",
		},
		ERROR: ShowServerResponseState{
			value: "error",
		},
		CLONING: ShowServerResponseState{
			value: "cloning",
		},
		TESTING: ShowServerResponseState{
			value: "testing",
		},
		FINISHED: ShowServerResponseState{
			value: "finished",
		},
	}
}

func (c ShowServerResponseState) Value() string {
	return c.value
}

func (c ShowServerResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServerResponseState) UnmarshalJSON(b []byte) error {
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

type ShowServerResponseFirmware struct {
	value string
}

type ShowServerResponseFirmwareEnum struct {
	BIOS ShowServerResponseFirmware
	UEFI ShowServerResponseFirmware
}

func GetShowServerResponseFirmwareEnum() ShowServerResponseFirmwareEnum {
	return ShowServerResponseFirmwareEnum{
		BIOS: ShowServerResponseFirmware{
			value: "BIOS",
		},
		UEFI: ShowServerResponseFirmware{
			value: "UEFI",
		},
	}
}

func (c ShowServerResponseFirmware) Value() string {
	return c.value
}

func (c ShowServerResponseFirmware) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServerResponseFirmware) UnmarshalJSON(b []byte) error {
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

type ShowServerResponseMigrationCycle struct {
	value string
}

type ShowServerResponseMigrationCycleEnum struct {
	CUTOVERING  ShowServerResponseMigrationCycle
	CUTOVERED   ShowServerResponseMigrationCycle
	CHECKING    ShowServerResponseMigrationCycle
	SETTING     ShowServerResponseMigrationCycle
	REPLICATING ShowServerResponseMigrationCycle
	SYNCING     ShowServerResponseMigrationCycle
}

func GetShowServerResponseMigrationCycleEnum() ShowServerResponseMigrationCycleEnum {
	return ShowServerResponseMigrationCycleEnum{
		CUTOVERING: ShowServerResponseMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: ShowServerResponseMigrationCycle{
			value: "cutovered",
		},
		CHECKING: ShowServerResponseMigrationCycle{
			value: "checking",
		},
		SETTING: ShowServerResponseMigrationCycle{
			value: "setting",
		},
		REPLICATING: ShowServerResponseMigrationCycle{
			value: "replicating",
		},
		SYNCING: ShowServerResponseMigrationCycle{
			value: "syncing",
		},
	}
}

func (c ShowServerResponseMigrationCycle) Value() string {
	return c.value
}

func (c ShowServerResponseMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServerResponseMigrationCycle) UnmarshalJSON(b []byte) error {
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
