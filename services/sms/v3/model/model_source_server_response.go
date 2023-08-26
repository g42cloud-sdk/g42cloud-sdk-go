package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type SourceServerResponse struct {
	Id *string `json:"id,omitempty"`

	Ip string `json:"ip"`

	Name string `json:"name"`

	OsType SourceServerResponseOsType `json:"os_type"`

	OsVersion *string `json:"os_version,omitempty"`

	OemSystem *bool `json:"oem_system,omitempty"`

	State *SourceServerResponseState `json:"state,omitempty"`

	MigrationCycle *SourceServerResponseMigrationCycle `json:"migration_cycle,omitempty"`
}

func (o SourceServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceServerResponse struct{}"
	}

	return strings.Join([]string{"SourceServerResponse", string(data)}, " ")
}

type SourceServerResponseOsType struct {
	value string
}

type SourceServerResponseOsTypeEnum struct {
	WINDOWS SourceServerResponseOsType
	LINUX   SourceServerResponseOsType
}

func GetSourceServerResponseOsTypeEnum() SourceServerResponseOsTypeEnum {
	return SourceServerResponseOsTypeEnum{
		WINDOWS: SourceServerResponseOsType{
			value: "WINDOWS",
		},
		LINUX: SourceServerResponseOsType{
			value: "LINUX",
		},
	}
}

func (c SourceServerResponseOsType) Value() string {
	return c.value
}

func (c SourceServerResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerResponseOsType) UnmarshalJSON(b []byte) error {
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

type SourceServerResponseState struct {
	value string
}

type SourceServerResponseStateEnum struct {
	UNAVAILABLE SourceServerResponseState
	WAITING     SourceServerResponseState
	INITIALIZE  SourceServerResponseState
	REPLICATE   SourceServerResponseState
	SYNCING     SourceServerResponseState
	STOPPING    SourceServerResponseState
	STOPPED     SourceServerResponseState
	DELETING    SourceServerResponseState
	ERROR       SourceServerResponseState
	CLONING     SourceServerResponseState
	TESTING     SourceServerResponseState
	FINISHED    SourceServerResponseState
}

func GetSourceServerResponseStateEnum() SourceServerResponseStateEnum {
	return SourceServerResponseStateEnum{
		UNAVAILABLE: SourceServerResponseState{
			value: "unavailable",
		},
		WAITING: SourceServerResponseState{
			value: "waiting",
		},
		INITIALIZE: SourceServerResponseState{
			value: "initialize",
		},
		REPLICATE: SourceServerResponseState{
			value: "replicate",
		},
		SYNCING: SourceServerResponseState{
			value: "syncing",
		},
		STOPPING: SourceServerResponseState{
			value: "stopping",
		},
		STOPPED: SourceServerResponseState{
			value: "stopped",
		},
		DELETING: SourceServerResponseState{
			value: "deleting",
		},
		ERROR: SourceServerResponseState{
			value: "error",
		},
		CLONING: SourceServerResponseState{
			value: "cloning",
		},
		TESTING: SourceServerResponseState{
			value: "testing",
		},
		FINISHED: SourceServerResponseState{
			value: "finished",
		},
	}
}

func (c SourceServerResponseState) Value() string {
	return c.value
}

func (c SourceServerResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerResponseState) UnmarshalJSON(b []byte) error {
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

type SourceServerResponseMigrationCycle struct {
	value string
}

type SourceServerResponseMigrationCycleEnum struct {
	CUTOVERING  SourceServerResponseMigrationCycle
	CUTOVERED   SourceServerResponseMigrationCycle
	CHECKING    SourceServerResponseMigrationCycle
	SETTING     SourceServerResponseMigrationCycle
	REPLICATING SourceServerResponseMigrationCycle
	SYNCING     SourceServerResponseMigrationCycle
}

func GetSourceServerResponseMigrationCycleEnum() SourceServerResponseMigrationCycleEnum {
	return SourceServerResponseMigrationCycleEnum{
		CUTOVERING: SourceServerResponseMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: SourceServerResponseMigrationCycle{
			value: "cutovered",
		},
		CHECKING: SourceServerResponseMigrationCycle{
			value: "checking",
		},
		SETTING: SourceServerResponseMigrationCycle{
			value: "setting",
		},
		REPLICATING: SourceServerResponseMigrationCycle{
			value: "replicating",
		},
		SYNCING: SourceServerResponseMigrationCycle{
			value: "syncing",
		},
	}
}

func (c SourceServerResponseMigrationCycle) Value() string {
	return c.value
}

func (c SourceServerResponseMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerResponseMigrationCycle) UnmarshalJSON(b []byte) error {
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
